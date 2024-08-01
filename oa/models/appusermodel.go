package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AppUserModel = (*customAppUserModel)(nil)

type (
	// AppUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserModel.
	AppUserModel interface {
		appUserModel
		withSession(session sqlx.Session) AppUserModel
	}

	customAppUserModel struct {
		*defaultAppUserModel
	}
)

// NewAppUserModel returns a model for the database table.
func NewAppUserModel(conn sqlx.SqlConn) AppUserModel {
	return &customAppUserModel{
		defaultAppUserModel: newAppUserModel(conn),
	}
}

func (m *customAppUserModel) withSession(session sqlx.Session) AppUserModel {
	return NewAppUserModel(sqlx.NewSqlConnFromSession(session))
}
