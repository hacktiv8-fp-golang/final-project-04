package main

import (
	"final-project-04/internal/database"
	"final-project-04/internal/router"
)

func main() {
	database.StartDB()
	router.StartServer()
}
