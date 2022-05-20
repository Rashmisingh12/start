package main

import (
	"Project1/database"
	Logger "Project1/logger"
	"Project1/route"
	"net/http"
)

func main() {
	database.Start()

	r := route.Router()
	Logger.CommonLog.Println("server is running")
	http.ListenAndServe(":9090", r)

}
