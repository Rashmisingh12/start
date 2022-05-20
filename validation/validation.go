package validation

import (
	"regexp"
	
)

var rxEmail = regexp.MustCompile(`^[A-Z]{20}.[a-z]{20}$`)
var first_name=regexp.MustCompile(`^[A-Z]{20}.[a-z]{20}$`)
var password=regexp.MustCompile(`^[A-Z]{20}.[a-z]{20}$`)

type message struct{
	first_name string
	Email string
	password string
	content string
	Message string
	Errors map[string]string
}
func (msg *message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}


	return len(msg.Errors) == 0
}

func (msg *message) Validation1() bool{
	msg.Errors=make(map[string]string)
	match:=first_name.Match([]byte(msg.first_name))
	if match == false{
		msg.Errors["firstname"]="please enter your firstname"
	}
	return len(msg.Errors) == 0
}

func (msg *message) Validation2() bool{
	msg.Errors=make(map[string]string)
    match:=password.Match([]byte(msg.password))
	if match == false{
       msg.Errors["password"]="please enter your password"
	}
	return len(msg.Errors) == 0
}