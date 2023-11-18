package router

import (
	"final-project-04/internal/controller"
	"final-project-04/internal/middleware"

	_ "final-project-04/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var PORT = ":8089"

// @Title Toko Belanja API
// @version 1.0
// @description API for an e-commerce platform
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email sobercoder@swagger.io
// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

	router.Run(PORT)
}
