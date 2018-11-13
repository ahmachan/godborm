package models

import (
	"database/sql"
	"errors"
	"fmt"

	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type DataModel struct {
	db        *sql.DB
	pk        string
	tablename string
	columnstr string
	where     string
	orderby   string
	limit     string
	param     []string
}

func (m *DataModel) Insert(params map[string]interface{}) (num int, err error) {
	if m.db == nil {
		return 0, errors.New("mysql not connect")
	}
	var keys []string
	var values []string
	if len(m.pk) != 0 {
		delete(params, m.pk)
	}
	for key, value := range params {
		keys = append(keys, key)
		switch value.(type) {
		case int, int64, int32:
			values = append(values, strconv.Itoa(value.(int)))
		case string:
			values = append(values, value.(string))
		case float32, float64:
			values = append(values, strconv.FormatFloat(value.(float64), 'f', -1, 64))

		}

	}
	fileValue := "'" + strings.Join(values, "','") + "'"
	fileds := "`" + strings.Join(keys, "`,`") + "`"
	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", m.tablename, fileds, fileValue)
	result, err := m.db.Exec(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("inster sql failure")
		return 0, err
	}
	newId, err := result.LastInsertId()
	lastId, _ := strconv.Atoi(strconv.FormatInt(newId, 10))
	if err != nil {
		err = errors.New("insert failure")
	}
	return lastId, err

}
