package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Creates a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param model.Category body model.CategoryCreate true "Category object to be created"
// @Success 201 {object} model.User "Category created successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /categories [post]
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

	context.JSON(http.StatusCreated, gin.H{
		"id":                  categoryResponse.ID,
		"type":                categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"created_at":          categoryResponse.CreatedAt,
	})
}

// GetAllCategories godoc
// @Summary Get All Categories.
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} model.Category "Categories fetched successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /categories [get]
func GetAllCategories(context *gin.Context) {
	categories, err := service.CategoryService.GetAllCategories()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary Update a Category.
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param categoryId path int true "Category ID"
// @Param model.Category body model.CategoryUpdate true "Category object to be updated"
// @Success 200 {object} model.User "Category updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /categories/{categoryId} [put]
func UpdateCategory(context *gin.Context) {
	var categoryUpdated model.CategoryUpdate

	if err := context.ShouldBindJSON(&categoryUpdated); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	id, _ := helper.GetIdParam(context, "categoryId")

	categoryResponse, err := service.CategoryService.UpdateCategory(&categoryUpdated, id)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":                  categoryResponse.ID,
		"type":                categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"updated_at":          categoryResponse.UpdatedAt,
	})
}

// DeleteCategory godoc
// @Summary Delete a Category.
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param categoryId path int true "Category ID"
// @Success 200 {object} model.Category "Category deleted successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /categories/{categoryId} [delete]
func DeleteCategory(context *gin.Context) {
	categoryId, _ := helper.GetIdParam(context, "categoryId")

	err := service.CategoryService.DeleteCategory(categoryId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}
