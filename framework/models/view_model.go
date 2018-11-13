package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	//"github.com/widuu/goini"
	"godborm/framework/db"
	"regexp"
	"strconv"
	"strings"
)

type ViewModel struct {
	db        *sql.DB
	tablename string
	param     []string
	columnstr string
	where     string
	pk        string
	orderby   string
	limit     string
	viewmodel map[string]map[string]interface{}
}

var ViewModelInstance map[string]map[string]interface{}

func InstanceViewModel(dbhandle *db.DbHandle) (*ViewModel, error) {
	c := new(ViewModel)
	c.db = dbhandle.Db
	return c, nil
}

func (m *ViewModel) FindViewModel() map[int]map[string]string {

	result := make(map[int]map[string]string)
	if m.db == nil {
		fmt.Printf("mysql not connect")
		return result
	}
	if len(m.param) == 0 {
		m.columnstr = "*"
	} else {
		if len(m.param) == 1 {
			m.columnstr = m.param[0]
		} else {
			m.columnstr = strings.Join(m.param, ",")
		}
	}

	queryVmStr := ""
	var columnFields []string
	my_viewmodel := m.viewmodel

	for keyA, valA := range my_viewmodel {
		aliseTableName := keyA
		tureTableName := ""
		joinTypeStr := "LEFT JOIN"
		onRelationStr := ""
		//fmt.Printf("keyA:%v \n",keyA)
		for subKeyA, subValA := range valA {
			switch subKeyA {
			case "_table":
				tureTableName = subValA.(string)
			case "_type":
				joinType := subValA.(string)
				if joinType == "LEFT" {
					joinTypeStr = "LEFT JOIN"
				}
			case "_on":
				onRelation := subValA.(string)
				onRelationStr = onRelation
			default:
				columnFields = append(columnFields, fmt.Sprintf("%v.%v AS %v", aliseTableName, escapeString(subKeyA), subValA))
			}

			//fmt.Printf("key,val<==>%v,%v \n",subKeyA, subValA)
			// subString = fmt.Sprintf("%v AS %v %v %v AS %v ON %v", m.columnstr, m_viewmodel, m.join, m.where, m.orderby, m.limit)
		}

		if onRelationStr != "" {
			queryVmStr = queryVmStr + " " + fmt.Sprintf("%v %v AS %v", joinTypeStr, tureTableName, aliseTableName)
			queryVmStr = queryVmStr + " ON " + onRelationStr
		} else {
			queryVmStr = queryVmStr + fmt.Sprintf("%v AS %v", tureTableName, aliseTableName)
			queryVmStr = queryVmStr + onRelationStr
		}

	}

	//println(queryVmStr)
	columnFieldStr := strings.Join(columnFields, ",")
	m.columnstr = columnFieldStr
	//t_user as user
	//left join t_user_captcha as captcha
	//ON user.id=captcha.user_id

	var m_viewmodel = queryVmStr //"t_user as user left join t_user_captcha as captcha ON user.id=captcha.user_id"
	//var m_viewmodel=m.viewmodel
	query := fmt.Sprintf("SELECT %v FROM %v %v %v %v", m.columnstr, m_viewmodel, m.where, m.orderby, m.limit)
	println(query)
	rows, err := m.db.Query(query)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("select sql failure")
	}
	result = QueryResult(rows)
	return result
}

func (m *ViewModel) FindAll() map[int]map[string]string {

	result := make(map[int]map[string]string)
	if m.db == nil {
		fmt.Printf("mysql not connect")
		return result
	}
	if len(m.param) == 0 {
		m.columnstr = "*"
	} else {
		if len(m.param) == 1 {
			m.columnstr = m.param[0]
		} else {
			m.columnstr = strings.Join(m.param, ",")
		}

	}

	query := fmt.Sprintf("Select %v from %v %v %v %v", m.columnstr, m.tablename, m.where, m.orderby, m.limit)
	rows, err := m.db.Query(query)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("select sql failure")
	}
	result = QueryResult(rows)
	return result
}

func (m *ViewModel) FindOne() map[int]map[string]string {
	empty := make(map[int]map[string]string)
	if m.db != nil {
		data := m.Limit(1).FindAll()
		return data
	}
	fmt.Printf("mysql not connect\r\n")
	return empty
}

func (m *ViewModel) Insert(param map[string]interface{}) (num int, err error) {
	if m.db == nil {
		return 0, errors.New("mysql not connect")
	}
	var keys []string
	var values []string
	if len(m.pk) != 0 {
		delete(param, m.pk)
	}
	for key, value := range param {
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
	i, err := result.LastInsertId()
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))
	if err != nil {
		err = errors.New("insert failure")
	}
	return s, err

}

func (m *ViewModel) Fileds(param ...string) *ViewModel {
	m.param = param
	return m
}

func (m *ViewModel) Update(param map[string]interface{}) (num int, err error) {
	if m.db == nil {
		return 0, errors.New("mysql not connect")
	}
	var setValue []string
	for key, value := range param {
		switch value.(type) {
		case int, int64, int32:
			set := fmt.Sprintf("%v = %v", key, value.(int))
			setValue = append(setValue, set)
		case string:
			set := fmt.Sprintf("%v = '%v'", key, value.(string))
			setValue = append(setValue, set)
		case float32, float64:
			set := fmt.Sprintf("%v = '%v'", key, strconv.FormatFloat(value.(float64), 'f', -1, 64))
			setValue = append(setValue, set)
		}

	}
	setData := strings.Join(setValue, ",")
	sql := fmt.Sprintf("UPDATE %v SET %v %v", m.tablename, setData, m.where)
	result, err := m.db.Exec(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("update sql failure")
		return 0, err
	}
	i, err := result.RowsAffected()
	if err != nil {
		err = errors.New("update failure")
		return 0, err
	}
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))

	return s, err
}

func (m *ViewModel) Delete(param string) (num int, err error) {
	if m.db == nil {
		return 0, errors.New("mysql not connect")
	}
	h := m.Where(param).FindOne()
	if len(h) == 0 {
		return 0, errors.New("no Value")
	}
	sql := fmt.Sprintf("DELETE FROM %v WHERE %v", m.tablename, param)
	result, err := m.db.Exec(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("delete sql failure")
		return 0, err
	}
	i, err := result.RowsAffected()
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))
	if i == 0 {
		err = errors.New("delete failure")
	}

	return s, err
}

func (m *ViewModel) Query(sql string) interface{} {
	if m.db == nil {
		return errors.New("mysql not connect")
	}
	var query = strings.TrimSpace(sql)
	s, err := regexp.MatchString(`(?i)^select`, query)
	if err == nil && s == true {
		result, _ := m.db.Query(sql)
		c := QueryResult(result)
		return c
	}
	exec, err := regexp.MatchString(`(?i)^(update|delete)`, query)
	if err == nil && exec == true {
		m_exec, err := m.db.Exec(query)
		if err != nil {
			return err
		}
		num, _ := m_exec.RowsAffected()
		id := strconv.FormatInt(num, 10)
		return id
	}

	insert, err := regexp.MatchString(`(?i)^insert`, query)
	if err == nil && insert == true {
		m_exec, err := m.db.Exec(query)
		if err != nil {
			return err
		}
		num, _ := m_exec.LastInsertId()
		id := strconv.FormatInt(num, 10)
		return id
	}
	result, _ := m.db.Exec(query)

	return result

}

func QueryResult(rows *sql.Rows) map[int]map[string]string {
	var result = make(map[int]map[string]string)
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanargs := make([]interface{}, len(values))
	for i := range values {
		scanargs[i] = &values[i]
	}

	var n = 1
	for rows.Next() {
		result[n] = make(map[string]string)
		err := rows.Scan(scanargs...)

		if err != nil {
			fmt.Println(err)
		}

		for i, v := range values {
			result[n][columns[i]] = string(v)
		}
		n++
	}

	return result
}

func (m *ViewModel) IocViewModel(m_viewmodel map[string]map[string]interface{}) *ViewModel {
	m.viewmodel = m_viewmodel
	return m
}

func (m *ViewModel) ViewModelIoc() *ViewModel {
	m.viewmodel = ViewModelInstance
	return m
}

func (m *ViewModel) Where(param string) *ViewModel {
	/*
		var pack_params []string
		for pkey, pval := range param {
			temp_field := fmt.Sprintf("%v = %v", pkey, escapeString(pval))
			append(pack_params, temp_field)
		}
		last_params := string.Join(pack_params, " and ")
	*/
	m.where = fmt.Sprintf("WHERE %v", param)
	return m
}

func (m *ViewModel) SetPk(pk string) *ViewModel {
	m.pk = pk
	return m
}

func (m *ViewModel) OrderBy(param string) *ViewModel {
	m.orderby = fmt.Sprintf("ORDER BY %v", param)
	return m
}

func (m *ViewModel) Limit(size ...int) *ViewModel {
	var end int
	start := size[0]
	if len(size) > 1 {
		end = size[1]
		m.limit = fmt.Sprintf("LIMIT %d,%d", start, end)
		return m
	}
	m.limit = fmt.Sprintf("LIMIT %d", start)
	return m
}

func escapeString(field string) string {
	return fmt.Sprintf("`%v`", field)
}

func facade() string {
	return ""
}

//the function will use friendly way to print the data
func Print(slice map[int]map[string]string) {
	for _, v := range slice {
		for key, value := range v {
			fmt.Println(key, value)
		}
		fmt.Println("---------------")
	}
}

func (m *ViewModel) DbClose() {
	m.db.Close()
}
