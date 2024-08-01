package svc

import (
	"oa/internal/config"
	"oa/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
	_conn  sqlx.SqlConn
}

func (s *ServiceContext) Sqlx() sqlx.SqlConn {
	if s._conn == nil {
		s._conn = sqlx.NewMysql(s.Config.DB)
	}
	return s._conn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
