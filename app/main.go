package main

import (
	"app/dbconfig"
	"app/handlers"
	"app/models"
	"fmt"
	"log"
	"net/http"
)

// def: main
func main() {
	// start server
	StartServer()
}

// func: start server
func StartServer() {
	dbconnstr := fmt.Sprintf("host=%s dbname=%s user=%s "+
		"password=%s port=%d sslmode=disable",
		dbconfig.HOST, dbconfig.DBNAME, dbconfig.USER, dbconfig.PASSWORD, dbconfig.PORT)
	models.InitDB(dbconnstr)

	http.HandleFunc("/users", handlers.AllUsersHandler)
	http.HandleFunc("/status", handlers.ServerStatusHandler)
	log.Println("Now listening...")
	http.ListenAndServe(":8000", nil)
}
