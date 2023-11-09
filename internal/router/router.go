package router

import "github.com/gin-gonic/gin"

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	router.Run(PORT)
}
