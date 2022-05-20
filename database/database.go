package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func GetDB() *sql.DB {
	return Db
}

func Start() {
	Db, err = sql.Open("mysql", "root:rashmi@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err)

	}
}
