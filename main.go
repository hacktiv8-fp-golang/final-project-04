package main

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/router"
)

// @title Toko Belanja API
// @version 1.0
// @description API for an e-commerce platform
// @host final-project-04-production.up.railway.app
// @BasePath /

func main() {
	database.StartDB()
	router.StartServer()
}
