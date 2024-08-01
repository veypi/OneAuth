// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	resourceFieldNames          = builder.RawFieldNames(&Resource{})
	resourceRows                = strings.Join(resourceFieldNames, ",")
	resourceRowsExpectAutoSet   = strings.Join(stringx.Remove(resourceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	resourceRowsWithPlaceHolder = strings.Join(stringx.Remove(resourceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOaResourceIdPrefix        = "cache:oa:resource:id:"
	cacheOaResourceAppIdNamePrefix = "cache:oa:resource:appId:name:"
)

type (
	resourceModel interface {
		Insert(ctx context.Context, data *Resource) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Resource, error)
		FindOneByAppIdName(ctx context.Context, appId string, name string) (*Resource, error)
		Update(ctx context.Context, data *Resource) error
		Delete(ctx context.Context, id int64) error
	}

	defaultResourceModel struct {
		sqlc.CachedConn
		table string
	}

	Resource struct {
		Id      int64          `db:"id"`
		Created time.Time      `db:"created"`
		Updated time.Time      `db:"updated"`
		AppId   string         `db:"app_id"`
		Name    string         `db:"name"`
		Des     sql.NullString `db:"des"`
	}
)

func newResourceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultResourceModel {
	return &defaultResourceModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`resource`",
	}
}

func (m *defaultResourceModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	oaResourceAppIdNameKey := fmt.Sprintf("%s%v:%v", cacheOaResourceAppIdNamePrefix, data.AppId, data.Name)
	oaResourceIdKey := fmt.Sprintf("%s%v", cacheOaResourceIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, oaResourceAppIdNameKey, oaResourceIdKey)
	return err
}

func (m *defaultResourceModel) FindOne(ctx context.Context, id int64) (*Resource, error) {
	oaResourceIdKey := fmt.Sprintf("%s%v", cacheOaResourceIdPrefix, id)
	var resp Resource
	err := m.QueryRowCtx(ctx, &resp, oaResourceIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultResourceModel) FindOneByAppIdName(ctx context.Context, appId string, name string) (*Resource, error) {
	oaResourceAppIdNameKey := fmt.Sprintf("%s%v:%v", cacheOaResourceAppIdNamePrefix, appId, name)
	var resp Resource
	err := m.QueryRowIndexCtx(ctx, &resp, oaResourceAppIdNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `app_id` = ? and `name` = ? limit 1", resourceRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, appId, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultResourceModel) Insert(ctx context.Context, data *Resource) (sql.Result, error) {
	oaResourceAppIdNameKey := fmt.Sprintf("%s%v:%v", cacheOaResourceAppIdNamePrefix, data.AppId, data.Name)
	oaResourceIdKey := fmt.Sprintf("%s%v", cacheOaResourceIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, resourceRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Created, data.Updated, data.AppId, data.Name, data.Des)
	}, oaResourceAppIdNameKey, oaResourceIdKey)
	return ret, err
}

func (m *defaultResourceModel) Update(ctx context.Context, newData *Resource) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	oaResourceAppIdNameKey := fmt.Sprintf("%s%v:%v", cacheOaResourceAppIdNamePrefix, data.AppId, data.Name)
	oaResourceIdKey := fmt.Sprintf("%s%v", cacheOaResourceIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, resourceRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Created, newData.Updated, newData.AppId, newData.Name, newData.Des, newData.Id)
	}, oaResourceAppIdNameKey, oaResourceIdKey)
	return err
}

func (m *defaultResourceModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOaResourceIdPrefix, primary)
}

func (m *defaultResourceModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultResourceModel) tableName() string {
	return m.table
}
