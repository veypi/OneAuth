package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ResourceModel = (*customResourceModel)(nil)

type (
	// ResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customResourceModel.
	ResourceModel interface {
		resourceModel
		withSession(session sqlx.Session) ResourceModel
	}

	customResourceModel struct {
		*defaultResourceModel
	}
)

// NewResourceModel returns a model for the database table.
func NewResourceModel(conn sqlx.SqlConn) ResourceModel {
	return &customResourceModel{
		defaultResourceModel: newResourceModel(conn),
	}
}

func (m *customResourceModel) withSession(session sqlx.Session) ResourceModel {
	return NewResourceModel(sqlx.NewSqlConnFromSession(session))
}
