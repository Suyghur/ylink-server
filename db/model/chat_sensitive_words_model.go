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
	chatSensitiveWordsFieldNames          = builder.RawFieldNames(&ChatSensitiveWords{})
	chatSensitiveWordsRows                = strings.Join(chatSensitiveWordsFieldNames, ",")
	chatSensitiveWordsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatSensitiveWordsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	chatSensitiveWordsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatSensitiveWordsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ChatSensitiveWordsModel interface {
		Insert(data *ChatSensitiveWords) (sql.Result, error)
		FindOne(id int64) (*ChatSensitiveWords, error)
		FindOneByWord(word string) (*ChatSensitiveWords, error)
		Update(data *ChatSensitiveWords) error
		Delete(id int64) error
		FindAll() (*[]ChatSensitiveWords, error)
	}

	defaultChatSensitiveWordsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ChatSensitiveWords struct {
		Id      int64  `db:"id"`
		Word    string `db:"word"`
		AddType int64  `db:"add_type"`
	}
)

func NewChatSensitiveWordsModel(conn sqlx.SqlConn) ChatSensitiveWordsModel {
	return &defaultChatSensitiveWordsModel{
		conn:  conn,
		table: "`chat_sensitive_words`",
	}
}

func (m *defaultChatSensitiveWordsModel) Insert(data *ChatSensitiveWords) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, chatSensitiveWordsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Word, data.AddType)
	return ret, err
}

func (m *defaultChatSensitiveWordsModel) FindOne(id int64) (*ChatSensitiveWords, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatSensitiveWordsRows, m.table)
	var resp ChatSensitiveWords
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

func (m *defaultChatSensitiveWordsModel) FindOneByWord(word string) (*ChatSensitiveWords, error) {
	var resp ChatSensitiveWords
	query := fmt.Sprintf("select %s from %s where `word` = ? limit 1", chatSensitiveWordsRows, m.table)
	err := m.conn.QueryRow(&resp, query, word)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatSensitiveWordsModel) Update(data *ChatSensitiveWords) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatSensitiveWordsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Word, data.AddType, data.Id)
	return err
}

func (m *defaultChatSensitiveWordsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultChatSensitiveWordsModel) FindAll() (*[]ChatSensitiveWords, error) {
	words := make([]ChatSensitiveWords, 0)
	querySql := fmt.Sprintf("select %s from %s", chatSensitiveWordsRows, m.table)
	err := m.conn.QueryRows(&words, querySql)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &words, err
}
