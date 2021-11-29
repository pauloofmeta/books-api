package main

import (
	"books-api/database"
	"books-api/server"
)

func main()  {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}