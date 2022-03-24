package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	chatConfigsFieldNames          = builder.RawFieldNames(&ChatConfigs{})
	chatConfigsRows                = strings.Join(chatConfigsFieldNames, ",")
	chatConfigsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatConfigsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	chatConfigsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatConfigsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ChatConfigsModel interface {
		Insert(data *ChatConfigs) (sql.Result, error)
		FindOne(id int64) (*ChatConfigs, error)
		FindOneByConfKey(confKey string) (*ChatConfigs, error)
		Update(data *ChatConfigs) error
		Delete(id int64) error
		FindAll() ([]ChatConfigs, error)
	}

	defaultChatConfigsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ChatConfigs struct {
		Id        int64  `db:"id"`
		ConfName  string `db:"conf_name"`
		ConfKey   string `db:"conf_key"`
		ConfValue string `db:"conf_value"`
	}
)

func NewChatConfigsModel(conn sqlx.SqlConn) ChatConfigsModel {
	return &defaultChatConfigsModel{
		conn:  conn,
		table: "`chat_configs`",
	}
}

func (m *defaultChatConfigsModel) Insert(data *ChatConfigs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, chatConfigsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ConfName, data.ConfKey, data.ConfValue)
	return ret, err
}

func (m *defaultChatConfigsModel) FindOne(id int64) (*ChatConfigs, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatConfigsRows, m.table)
	var resp ChatConfigs
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatConfigsModel) FindOneByConfKey(confKey string) (*ChatConfigs, error) {
	var resp ChatConfigs
	query := fmt.Sprintf("select %s from %s where `conf_key` = ? limit 1", chatConfigsRows, m.table)
	err := m.conn.QueryRow(&resp, query, confKey)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatConfigsModel) Update(data *ChatConfigs) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatConfigsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ConfName, data.ConfKey, data.ConfValue, data.Id)
	return err
}

func (m *defaultChatConfigsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultChatConfigsModel) FindAll() ([]ChatConfigs, error) {
	configs := make([]ChatConfigs, 0)
	querySql := fmt.Sprintf("select %s from %s", configsRows, m.table)
	err := m.conn.QueryRows(&configs, querySql)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return configs, err
}
