// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tokenFieldNames          = builder.RawFieldNames(&Token{})
	tokenRows                = strings.Join(tokenFieldNames, ",")
	tokenRowsExpectAutoSet   = strings.Join(stringx.Remove(tokenFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tokenRowsWithPlaceHolder = strings.Join(stringx.Remove(tokenFieldNames, "`code`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tokenModel interface {
		Insert(ctx context.Context, data *Token) (sql.Result, error)
		FindOne(ctx context.Context, code string) (*Token, error)
		Update(ctx context.Context, data *Token) error
		Delete(ctx context.Context, code string) error
	}

	defaultTokenModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Token struct {
		Code     string    `db:"code"`
		Created  time.Time `db:"created"`
		Updated  time.Time `db:"updated"`
		Expired  time.Time `db:"expired"`
		ClientId string    `db:"client_id"`
		AppId    string    `db:"app_id"`
		UserId   string    `db:"user_id"`
		Meta     string    `db:"meta"`
	}
)

func newTokenModel(conn sqlx.SqlConn) *defaultTokenModel {
	return &defaultTokenModel{
		conn:  conn,
		table: "`token`",
	}
}

func (m *defaultTokenModel) Delete(ctx context.Context, code string) error {
	query := fmt.Sprintf("delete from %s where `code` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, code)
	return err
}

func (m *defaultTokenModel) FindOne(ctx context.Context, code string) (*Token, error) {
	query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", tokenRows, m.table)
	var resp Token
	err := m.conn.QueryRowCtx(ctx, &resp, query, code)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTokenModel) Insert(ctx context.Context, data *Token) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, tokenRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Code, data.Created, data.Updated, data.Expired, data.ClientId, data.AppId, data.UserId, data.Meta)
	return ret, err
}

func (m *defaultTokenModel) Update(ctx context.Context, data *Token) error {
	query := fmt.Sprintf("update %s set %s where `code` = ?", m.table, tokenRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Created, data.Updated, data.Expired, data.ClientId, data.AppId, data.UserId, data.Meta, data.Code)
	return err
}

func (m *defaultTokenModel) tableName() string {
	return m.table
}
