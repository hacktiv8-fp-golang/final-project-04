package router

import (
	"final-project-04/internal/controller"
	"final-project-04/internal/middleware"

	"github.com/gin-gonic/gin"
	_ "final-project-04/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

var PORT = ":8080"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}
