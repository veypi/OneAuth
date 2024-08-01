package user

import (
	"context"
	"strings"
	"time"

	"oa/internal/svc"
	"oa/internal/types"
	"oa/models"

	"github.com/google/uuid"
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

	m := models.NewUserModel(l.svcCtx.Sqlx())
	u := &models.User{
		Id:       strings.ReplaceAll(uuid.New().String(), "-", ""),
		Created:  time.Now(),
		Updated:  time.Now(),
		Username: req.Username,
	}
	l.Infof("user: %v", u.Id)
	_, e := m.Insert(l.ctx, u)
	return e
}
