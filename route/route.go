package route

import (
	middleware "Project1/service"

	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/insert", middleware.CrossOrigin(middleware.Insert)).Methods("POST")
	r.HandleFunc("/getuser", middleware.CrossOrigin(middleware.GetUser)).Methods("GET")
	// r.HandleFunc("/update", middleware.CrossOrigin(middleware.Update)).Methods("PUT")
	r.HandleFunc("/getuserbyid/{id}", middleware.CrossOrigin(middleware.GetUserByID)).Methods("GET")
	r.HandleFunc("/search", middleware.CrossOrigin(middleware.Search)).Methods("GET")
	r.HandleFunc("/deleteuser", middleware.CrossOrigin(middleware.DeleteUser)).Methods("DELETE")

	r.Handle("/favicon.ico", http.NotFoundHandler())

	return r
}
