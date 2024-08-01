package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppModel = (*customAppModel)(nil)

type (
	// AppModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppModel.
	AppModel interface {
		appModel
	}

	customAppModel struct {
		*defaultAppModel
	}
)

// NewAppModel returns a model for the database table.
func NewAppModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AppModel {
	return &customAppModel{
		defaultAppModel: newAppModel(conn, c, opts...),
	}
}
