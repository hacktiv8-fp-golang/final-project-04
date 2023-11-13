package router

import (
	"final-project-04/internal/controller"
	"final-project-04/internal/middleware"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.Use(middleware.Authentication())
		userRouter.PATCH("/topup", controller.UpdateBalance)
	}

	categoriesRouter := router.Group("/categories")
	{
		categoriesRouter.Use(middleware.Authentication())
		categoriesRouter.POST("/",middleware.Authorization(), controller.CreateCategory)
		categoriesRouter.GET("/")
		categoriesRouter.PATCH("/:categoryId", middleware.Authorization(), controller.UpdateCategory)
		categoriesRouter.DELETE("/:categoryId")
	}

	productsRouter := router.Group("/products")
	{
		productsRouter.Use(middleware.Authentication())
		productsRouter.POST("/", middleware.Authorization(), controller.CreateProduct)
		productsRouter.GET("/")
		productsRouter.PUT("/:productId")
		productsRouter.DELETE("/:productId")
	}

	transactionHistoryRouter := router.Group("/transactions")
	{
		transactionHistoryRouter.POST("/")
		transactionHistoryRouter.GET("/my-transactions")
		transactionHistoryRouter.GET("/user-transactions")
	}

	router.Run(PORT)
}
