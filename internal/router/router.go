package router

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/controller"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/middleware"
	"os"

	_ "github.com/hacktiv8-fp-golang/final-project-04/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
		categoriesRouter.POST("/", middleware.AdminAuthorization(), controller.CreateCategory)
		categoriesRouter.GET("/", controller.GetAllCategories)
		categoriesRouter.PATCH("/:categoryId", middleware.AdminAuthorization(), middleware.CategoryAuthorization(), controller.UpdateCategory)
		categoriesRouter.DELETE("/:categoryId", middleware.AdminAuthorization(), middleware.CategoryAuthorization(), controller.DeleteCategory)
	}

	productsRouter := router.Group("/products")
	{
		productsRouter.Use(middleware.Authentication())
		productsRouter.POST("/", middleware.AdminAuthorization(), controller.CreateProduct)
		productsRouter.GET("/", controller.GetAllProducts)
		productsRouter.PUT("/:productId", middleware.AdminAuthorization(), middleware.ProductAuthorization(), controller.UpdateProduct)
		productsRouter.DELETE("/:productId", middleware.AdminAuthorization(), middleware.ProductAuthorization(), controller.DeleteProduct)
	}

	transactionHistoryRouter := router.Group("/transactions")
	{
		transactionHistoryRouter.Use(middleware.Authentication())
		transactionHistoryRouter.POST("/", controller.CreateTransaction)
		transactionHistoryRouter.GET("/my-transactions", controller.GetTransactionsByUserID)
		transactionHistoryRouter.GET("/user-transactions", middleware.AdminAuthorization(),controller.GetAllTransaction)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var PORT = os.Getenv("PORT")

	router.Run(":" +PORT)
}
