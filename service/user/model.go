package model

import(
	"encoding/json"
)

type User struct{
	ID        string `json:"id"`
    Lastname  string `json:"lastname"`
    Firstname string `json:"firstname"`
	Email string `json:"email"`
	Password string `json:"password"`
	 Dob      int   `json:"age"`
    
}

var user= User{}
err:=json.Unmarshall(User,&user)

err:= json.NewDecoder().Decode(&user);
if err!=nil{
	fmt.Println()
}