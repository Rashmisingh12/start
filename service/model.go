package service

import (
	
	"time"
)

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Dob       int    `json:"dob"`
	DeletedAt time.Time `json:"deletedate"`
}

type Users struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	Dob       string
}

type ReUser struct{
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
  Archived   bool    `json:"archived"`
	Sort       string `json:"sort"`
	Page   int  `json:"page"`
	Items int `json:"items"`
}
