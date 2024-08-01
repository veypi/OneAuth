package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TokenModel = (*customTokenModel)(nil)

type (
	// TokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokenModel.
	TokenModel interface {
		tokenModel
		withSession(session sqlx.Session) TokenModel
	}

	customTokenModel struct {
		*defaultTokenModel
	}
)

// NewTokenModel returns a model for the database table.
func NewTokenModel(conn sqlx.SqlConn) TokenModel {
	return &customTokenModel{
		defaultTokenModel: newTokenModel(conn),
	}
}

func (m *customTokenModel) withSession(session sqlx.Session) TokenModel {
	return NewTokenModel(sqlx.NewSqlConnFromSession(session))
}
