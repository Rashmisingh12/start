package db

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB
var err error
func init(){
	db, err = sql.Open("mysql", "root:rashmi@tcp(localhost:3306)/first?charset=utf8")
	if err != nil {
		fmt.Println(err)
		defer db.Close()
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)

	}
}

func create(w http.ResponseWriter,req *http.Request){


	stmt, err := db.Prepare(`create table employee1(userId int primary key,Fname varchar(20),Lname varchar(20),Email varchar(50))`)
	fmt.Println(err)

	defer stmt.Close()

	r, err := stmt.Exec()
	fmt.Println(err)

	n, err := r.RowsAffected()
	fmt.Println(err)
	fmt.Fprintln(w, "created table employee1", n)
}
