package controller

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(context *gin.Context) {
	var product model.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	productResponse, err := service.ProductService.CreateProduct(&product)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id": productResponse.ID,
		"title": productResponse.Title,
		"price": productResponse.Price,
		"stock": productResponse.Stock,
		"category_id": productResponse.CategoryID,
		"created_at": productResponse.CreatedAt,
	})
}
