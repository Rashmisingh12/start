package service

import (
	"Project1/database"
	// "Project1/model"

	// "database/sql"

	Logger "Project1/logger"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secretkey")

type Credentials struct {
	Password string `json:"password"`
	email    string `json:"email"`
}

type Claims struct {
	email string `json:"email"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	Db := database.GetDB()
	var query string
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query = fmt.Sprintf(`select password from Users where email='%s'`, creds.email)
	row, err := Db.Query(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var expectedPassword string
	for row.Next() {
		row.Scan(&expectedPassword)
	}

	if expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		email: creds.email,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Insert(w http.ResponseWriter, r *http.Request) {
	Db := database.GetDB()
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    var user User
    
		ID, _ := strconv.Atoi(user.ID)

		DOB, _ := time.Parse("2006-01-02",strconv.Itoa(user.Dob))
		bs, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		query := fmt.Sprintf(`INSERT INTO Users(user_id,first_name,last_name,email,password,dob,created_at,archived) VALUES(%d,"%s","%s","%s","%s","%v","%v","%d")`, ID, user.Firstname, user.Lastname,user.Email, string(bs), DOB.Format("2006-01-02"), time.Now(), 0)
		_, err = Db.Exec(query)
		if err != nil {
			log.Println(err)
			Logger.CommonLog.Println("Server is running")

		}
		json.NewDecoder(r.Body).Decode(&user)

		

	}



func GetUser(w http.ResponseWriter, r *http.Request) {
	Db := database.GetDB()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	selDB, err := Db.Query("SELECT user_id,first_name,last_name,email,dob FROM Users")
	if err != nil {
		panic(err.Error())
	}

	var res []Users
	for selDB.Next() {
		var res1 Users
		selDB.Scan(&res1.ID, &res1.Firstname, &res1.Lastname, &res1.Email, &res1.Dob)
		res = append(res, res1)
	}
	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	Db := database.GetDB()

	usr := r.URL.Query().Get("id")

	_, err := Db.Query(`SELECT * FROM Users WHERE userid=$? `, usr)
	if err != nil {
		fmt.Print(err)
		Logger.CommonLog.Println("server is running")
	}
	query := fmt.Sprintf(`update Users set archived = true where user_id = %s`, usr)
	_, err = Db.Exec(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("DELETE")

}

func GetUserByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	Db := database.GetDB()

	var user Users

	// create the select sql query
	sqlStatement := `SELECT user_id,first_name,last_name,email,dob FROM Users WHERE user_id=?`

	// execute the sql statement
	err := Db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Dob)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)

}

func Search(w http.ResponseWriter, req *http.Request) {
	fmt.Println("search started")
	var user []Users
	fmt.Println("line 180")
	Db := database.GetDB()
	fmt.Println("line 182")

	Query := getQuery(req)
	fmt.Println("line 185")

	fmt.Println(Query)
	row, err := Db.Query(Query)
	if err != nil {
		fmt.Println(err)
	}

	
	

	for row.Next() {
		var res1 Users
		
		row.Scan(&res1.ID, &res1.Firstname, &res1.Lastname, &res1.Email, &res1.Dob)
		user = append(user, res1)
	}
	

	json.NewEncoder(w).Encode(user)
}

func getQuery(req *http.Request) string {

	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	email := req.URL.Query().Get("email")
	sortby := req.URL.Query().Get("sortby")
	archived := req.URL.Query().Get("archived")
	order := req.URL.Query().Get("order")
	page := req.URL.Query().Get("page")
	items := req.URL.Query().Get("items")

	query := "select user_id,first_name,last_name,email,dob from Users"
	if archived == "true" {
		query = "select user_id,first_name,last_name,email,dob from Users"
	}
	if id != "" {
		query += " and user_id=" + id
	}
	if name != "" {
		query += ` and first_name like '%` +name+ `%' or last_name like '%` + name+ `%'`
	}
	if email != "" {
		query += ` and email like '%` +email+ `'%'`
	}
	if sortby != "" {
		if order != "" {
			query += ` order by` + sortby + `` + order
		} else {
			query += ` order by` + sortby + `` + order
		}
	}
	if items == "" {
		items = "3"
	}
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)
	i, _ := strconv.Atoi(items)
	query += fmt.Sprintf(` LIMIT %d OFFSET %d`, i, (p-1)*i)
	fmt.Println(query)
	return query
}
