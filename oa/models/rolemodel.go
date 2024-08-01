package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		withSession(session sqlx.Session) RoleModel
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn),
	}
}

func (m *customRoleModel) withSession(session sqlx.Session) RoleModel {
	return NewRoleModel(sqlx.NewSqlConnFromSession(session))
}
