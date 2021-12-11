package fs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/webdav"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/models/index"
	"github.com/veypi/utils/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func updateUserSize(uid uint, delta int64) error {
	if uid == 0 {
		return nil
	}
	err := cfg.DB().Model(&models.User{}).Where("ID = ?", uid).
		Update("Used", gorm.Expr("Used + ?", delta)).Error
	return err
}

func removeFile(r *http.Request, prefix string, h *webdav.Handler) error {
	f := &models.File{}
	f.OwnerID = h.OwnerID
	f.Path = prefix
	res, err := index.File.Get().Id(f.ID()).Do(r.Context())
	if err != nil {
		if elastic.IsNotFound(err) {
			return nil
		}
		return err
	}
	_, err = index.File.Delete().Id(f.ID()).Do(r.Context())
	if err == nil {
		err = json.Unmarshal(res.Source, f)
		if err != nil {
			return err
		}
		uid, _ := strconv.Atoi(h.OwnerID)
		if f.Size > 0 && uid > 0 {
			err = updateUserSize(uint(uid), -int64(f.Size))
		}
		_ = addHistory(r, h, models.ActDelete, res.Id, f.Path)
	}
	return err
}

func updateFile(ctx context.Context, fileID string, props interface{}) error {
	_, err := index.File.Update().Id(fileID).Doc(props).Do(ctx)
	return err
}

func getFile(ctx context.Context, path string, h *webdav.Handler) (*models.File, error) {
	f := &models.File{}
	f.OwnerID = h.OwnerID
	f.Path = path
	//body, err := cfg.ES().Update().Index(models.IndexFile).Id(f.ID()).Upsert(f).Do(ctx)
	body, err := index.File.Get().Id(f.ID()).Do(ctx)
	if elastic.IsNotFound(err) {
		f.CreatedAt = time.Now()
		_, err = index.File.Index().Id(f.ID()).BodyJson(f).Do(ctx)
		return f, err
	} else if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body.Source, f)
	return f, err
}

func addHistory(r *http.Request, h *webdav.Handler, act models.Action, fileID string, tag string) error {
	his := &models.History{
		Action:  act,
		OwnerID: h.OwnerID,
		ActorID: h.ActorID,
		FileID:  fileID,
		Tag:     tag,
		IP:      r.RemoteAddr,
	}
	his.CreatedAt = time.Now()
	_, err := index.History.Index().BodyJson(his).Do(r.Context())
	if err != nil {
		return err
	}
	if act == models.ActGet && fileID != "" {
		_, err = index.File.Update().Id(fileID).
			Script(elastic.NewScript("ctx._source.Count ++")).Do(r.Context())
	}
	return err
}

func CleanHistory(days uint) error {
	res, err := index.History.DeleteByQuery().
		Query(elastic.NewRangeQuery("CreatedAt").Lt(fmt.Sprintf("now-%dd/d", days))).
		Do(cfg.Ctx)
	if err == nil && res.Deleted > 0 {
		log.Info().Msgf("clean %d records of history.", res.Deleted)
	}
	return err
}
