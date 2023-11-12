package controller

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(context *gin.Context) {
	var category model.Category

	if err := context.ShouldBindJSON(&category); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	categoryResponse, err := service.CategoryService.CreateCategory(&category)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id": categoryResponse.ID,
		"type": categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"created_at": categoryResponse.CreatedAt,
	})
}
