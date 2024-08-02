package user

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"oa/errs"
	"oa/internal/svc"
	"oa/internal/types"
	"oa/models"

	"github.com/google/uuid"
	"github.com/veypi/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req *types.RegReq) error {
	// todo: add your logic here and delete this line

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := models.NewUserModel(l.svcCtx.Sqlx())
	u := &models.User{
		Id:       strings.ReplaceAll(uuid.New().String(), "-", ""),
		Created:  time.Now(),
		Updated:  time.Now(),
		Username: req.Username,
		RealCode: utils.RandSeq(32),
		Icon:     fmt.Sprintf("/media/icon/default/%04d.jpg", r.Intn(230)),
		Space:    300,
	}
	var err error
	u.CheckCode, err = utils.AesEncrypt(u.RealCode, []byte(req.Pwd))
	if err != nil {
		return errs.ArgsInvalid.WithErr(err)
	}
	l.Infof("user: %v", u.Id)
	_, err = m.Insert(l.ctx, u)
	return err
}
