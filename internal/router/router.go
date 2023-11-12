package router

import (
	"final-project-04/internal/controller"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
	}

	router.Run(PORT)
}
