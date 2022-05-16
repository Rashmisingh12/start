package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goriila/mux"
	"gorm.io/driver/postgres"
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

    r := mux.NewRouter()
    usersR := r.PathPrefix("/users").Subrouter()
    usersR.Path("").Methods(http.MethodGet).HandlerFunc(getAllUsers)
    usersR.Path("").Methods(http.MethodPost).HandlerFunc(createUser)
    usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getUserByID)
    usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(updateUser)
    usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deleteUser)
    fmt.Println("Start listening")
    fmt.Println(http.ListenAndServe(":8080", r))
}
// func getAllUsers(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("Not implemented")
// }
// func getUserByID(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("Not implemented")
// }
// func updateUser(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("Not implemented")
// }
// func deleteUser(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("Not implemented")
// }
// func createUser(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("Not implemented")
// }

func initDB() (*gorm.DB, error) {
	dataSourceName := "host=192.168.2.139 user=postgres password=1234 dbname=userdb port=5432"
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(&User{})

	return db, err
}

type userHandler struct {
	db *gorm.DB
}
func (uh userHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	if err := uh.db.Find(&users).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "Error on DB find for all users", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(w).Encode(users); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}
func (uh userHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user := User{ID: id}
	if err := uh.db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error on DB find for user with id: %s", id), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}
func (uh userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.db.Create(&u).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
