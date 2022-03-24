package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	configsFieldNames          = builderx.RawFieldNames(&Configs{})
	configsRows                = strings.Join(configsFieldNames, ",")
	configsRowsExpectAutoSet   = strings.Join(stringx.Remove(configsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	configsRowsWithPlaceHolder = strings.Join(stringx.Remove(configsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheConfigsIdPrefix      = "cache:configs:id:"
	cacheConfigsConfKeyPrefix = "cache:configs:confKey:"
)

type (
	ConfigsModel interface {
		Insert(data Configs) (sql.Result, error)
		FindOne(id int64) (*Configs, error)
		FindOneByConfKey(confKey string) (*Configs, error)
		Update(data Configs) error
		Delete(id int64) error
		FindAll() ([]Configs, error)
	}

	defaultConfigsModel struct {
		sqlc.CachedConn
		table string
	}

	Configs struct {
		Id        int64  `db:"id"`
		ConfName  string `db:"conf_name"`
		ConfKey   string `db:"conf_key"`
		ConfValue string `db:"conf_value"`
	}
)

func NewConfigsModel(conn sqlx.SqlConn, c cache.CacheConf) ConfigsModel {
	return &defaultConfigsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`configs`",
	}
}

func (m *defaultConfigsModel) Insert(data Configs) (sql.Result, error) {
	configsConfKeyKey := fmt.Sprintf("%s%v", cacheConfigsConfKeyPrefix, data.ConfKey)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, configsRowsExpectAutoSet)
		return conn.Exec(query, data.ConfName, data.ConfKey, data.ConfValue)
	}, configsConfKeyKey)
	return ret, err
}

func (m *defaultConfigsModel) FindOne(id int64) (*Configs, error) {
	configsIdKey := fmt.Sprintf("%s%v", cacheConfigsIdPrefix, id)
	var resp Configs
	err := m.QueryRow(&resp, configsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", configsRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultConfigsModel) FindOneByConfKey(confKey string) (*Configs, error) {
	configsConfKeyKey := fmt.Sprintf("%s%v", cacheConfigsConfKeyPrefix, confKey)
	var resp Configs
	err := m.QueryRowIndex(&resp, configsConfKeyKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `conf_key` = ? limit 1", configsRows, m.table)
		if err := conn.QueryRow(&resp, query, confKey); err != nil {
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

func (m *defaultConfigsModel) Update(data Configs) error {
	configsIdKey := fmt.Sprintf("%s%v", cacheConfigsIdPrefix, data.Id)
	configsConfKeyKey := fmt.Sprintf("%s%v", cacheConfigsConfKeyPrefix, data.ConfKey)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, configsRowsWithPlaceHolder)
		return conn.Exec(query, data.ConfName, data.ConfKey, data.ConfValue, data.Id)
	}, configsIdKey, configsConfKeyKey)
	return err
}

func (m *defaultConfigsModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	configsIdKey := fmt.Sprintf("%s%v", cacheConfigsIdPrefix, id)
	configsConfKeyKey := fmt.Sprintf("%s%v", cacheConfigsConfKeyPrefix, data.ConfKey)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, configsIdKey, configsConfKeyKey)
	return err
}

func (m *defaultConfigsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheConfigsIdPrefix, primary)
}

func (m *defaultConfigsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", configsRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultConfigsModel) FindAll() ([]Configs, error) {
	configs := make([]Configs, 0)
	querySql := fmt.Sprintf("select %s from %s", configsRows, m.table)
	err := m.CachedConn.QueryRowsNoCache(&configs, querySql)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return configs, err
}
