package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AppModel = (*customAppModel)(nil)

type (
	// AppModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppModel.
	AppModel interface {
		appModel
		withSession(session sqlx.Session) AppModel
	}

	customAppModel struct {
		*defaultAppModel
	}
)

// NewAppModel returns a model for the database table.
func NewAppModel(conn sqlx.SqlConn) AppModel {
	return &customAppModel{
		defaultAppModel: newAppModel(conn),
	}
}

func (m *customAppModel) withSession(session sqlx.Session) AppModel {
	return NewAppModel(sqlx.NewSqlConnFromSession(session))
}
