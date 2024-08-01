package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		withSession(session sqlx.Session) UserRoleModel
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn sqlx.SqlConn) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn),
	}
}

func (m *customUserRoleModel) withSession(session sqlx.Session) UserRoleModel {
	return NewUserRoleModel(sqlx.NewSqlConnFromSession(session))
}
