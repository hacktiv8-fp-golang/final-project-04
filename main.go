package main

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/router"
)

func main() {
	database.StartDB()
	router.StartServer()
}
