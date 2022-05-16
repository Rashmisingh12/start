package routes


import(
	"net/http"
	"fmt"
	"github.com/gorilla/mux"

)
func main(){
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

//       r := mux.NewRouter()

//     r.HandleFunc("/employee", employee).Methods("GET")
// 	r.HandleFunc("/create",create).Methods("POST")
// 	r.HandleFunc("/insert", insert).Methods("POST")
// 	r.HandleFunc("/read", read).Methods("GET")
// 	r.HandleFunc("/update",update).Methods("GET")
// 	r.HandleFunc("/delete",delete).Methods("DELETE")

// 	r.Handle("/favicon.ico", http.NotFoundHandler())
// 	fmt.Println("start listening")
//    http.ListenAndServe(":8081", r)
// }
// func index(w httfunc create(w http.ResponseWriter,req *http.Request){
    
// }
// p.ResponseWriter, req *http.Request) {
// 	_, err := io.WriteString(w, "succesfully completed")
// 	fmt.Println(err)


