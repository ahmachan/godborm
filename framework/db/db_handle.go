package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/widuu/goini"
)

type DbHandle struct {
	Db *sql.DB
}

//use goini read the configuration file and connect the mysql database
func SetConfig(filename string) (*DbHandle, error) {
	fmt.Println(filename)
	conf := goini.SetConfig(filename)
	charset := conf.GetValue("database", "charset")
	username := conf.GetValue("database", "username")
	password := conf.GetValue("database", "password")
	hostname := conf.GetValue("database", "hostname")
	database := conf.GetValue("database", "database")
	port := conf.GetValue("database", "port")
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+hostname+":"+port+")/"+database+"?charset="+charset)
	err = db.Ping()
	c := new(DbHandle)
	if err != nil {
		//if connect error then return the error message
		return c, err
	}

	c.Db = db

	return c, nil
}
