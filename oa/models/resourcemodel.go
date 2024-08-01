package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ResourceModel = (*customResourceModel)(nil)

type (
	// ResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customResourceModel.
	ResourceModel interface {
		resourceModel
	}

	customResourceModel struct {
		*defaultResourceModel
	}
)

// NewResourceModel returns a model for the database table.
func NewResourceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ResourceModel {
	return &customResourceModel{
		defaultResourceModel: newResourceModel(conn, c, opts...),
	}
}
