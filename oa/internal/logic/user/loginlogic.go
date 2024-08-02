package user

import (
	"context"
	"database/sql"
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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (string, error) {
	// todo: add your logic here and delete this line
	m := models.NewUserModel(l.svcCtx.Sqlx())
	var u *models.User
	var err error
	switch req.Typ {
	case "email":
		u, err = m.FindOneByEmail(l.ctx, sql.NullString{String: req.Id, Valid: true})
	case "phone":
		u, err = m.FindOneByPhone(l.ctx, sql.NullString{String: req.Id, Valid: true})
	default:
		u, err = m.FindOneByUsername(l.ctx, req.Id)
	}
	if err != nil {
		return "", errs.UserNotFound.WithErr(err)
	}
	temp, err := utils.AesDecrypt(u.CheckCode, []byte(req.Pwd))
	if err != nil || temp != u.RealCode {
		return "", errs.UserPwdInvalid
	}
	t := models.Token{
		Code:     strings.ReplaceAll(uuid.New().String(), "-", ""),
		Expired:  time.Now().Add(time.Hour * 24),
		ClientId: req.Client,
		AppId:    l.svcCtx.Config.UUID,
		UserId:   u.Id,
	}
	_, err = models.NewTokenModel(l.svcCtx.Sqlx()).Insert(l.ctx, &t)

	return "", err
}
