package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:rashmi@tcp(localhost:3306)/first?charset=utf8")
	if err != nil {
		fmt.Println(err)
		defer db.Close()
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)

	}
	http.HandleFunc("/", index)
	http.HandleFunc("/employee", employee)
	http.HandleFunc("/create",create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update",update)
	http.HandleFunc("/delete",delete)
	http.HandleFunc("/drop",drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "succesfully completed")
	fmt.Println(err)
}

func employee(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`select Ename from employee`)
	fmt.Println(err)
	defer rows.Close()
	var s, Ename string
	s = "Retrieve record\n"
	for rows.Next() {
		err = rows.Scan(&Ename)
		fmt.Println(err)
		s += Ename + "\n"
	}
	fmt.Fprint(w, s)

}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`create table employee1(Ename varchar(20),salary int)`)
	fmt.Println(err)

	defer stmt.Close()

	r, err := stmt.Exec()
	fmt.Println(err)

	n, err := r.RowsAffected()
	fmt.Println(err)
	fmt.Fprintln(w, "created table employee1", n)
}
func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT into employee1 values("Rashmi",25000),("Akash",30000),("Rahul",25000)`)
	fmt.Println(err)

	defer stmt.Close()

	r, err := stmt.Exec()
	fmt.Println(err)

	n, err := r.RowsAffected()
	fmt.Println(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}
func read(w http.ResponseWriter, req *http.Request) {
	stmt, _ := db.Query(`select * from employee1`)
	// fmt.Println(err)
	// defer stmt.Close()

	for stmt.Next() {
		var Ename string
		var saly int
		err = stmt.Scan(&Ename, &saly)
		fmt.Fprintln(w, "ham yaha hai", Ename, saly)
	}
	// fmt.Println(err)

	// n,err:=r.RowsAffected()
	// fmt.Fprintln(w,n)
}
func update(w http.ResponseWriter,req *http.Request){
	stmt,err:=db.Prepare(`update employee1 set Ename="Nisha" where Ename="Rashmi"`)
	fmt.Println(err)
	defer stmt.Close()

	r,err:=stmt.Exec()
	fmt.Println(err)

	n,err:=r.RowsAffected()
	fmt.Println(err)

	fmt.Fprintln(w,"UPDATED RECORD",n)
}
func delete(w http.ResponseWriter,req *http.Request){
	stmt,err:=db.Prepare(`delete from employee1 where Ename="Rahul" `)
	fmt.Println(err)

	defer stmt.Close()

	r,err:=stmt.Exec()
	fmt.Println(err)

	n,err:=r.RowsAffected()
	fmt.Println(err)

	fmt.Fprintln(w,"deleted row from employee1",n)
}

func drop(w http.ResponseWriter,req *http.Request){
	stmt,err:=db.Prepare(`drop table employee`)
	fmt.Println(err)
	defer stmt.Close()

	_,err=stmt.Exec()
	fmt.Println(err)

	fmt.Fprintln(w,"droped table employee1")
}

