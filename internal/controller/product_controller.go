package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Creates a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param model.Product body model.ProductCreate true "Product object to be created"
// @Success 201 {object} model.User "Product created successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /products [post]
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
		"id":          productResponse.ID,
		"title":       productResponse.Title,
		"price":       productResponse.Price,
		"stock":       productResponse.Stock,
		"category_id": productResponse.CategoryID,
		"created_at":  productResponse.CreatedAt,
	})
}

// GetAllProducts godoc
// @Summary Get All Products.
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} model.Product "Products fetched successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /products [get]
func GetAllProducts(context *gin.Context) {
	products, err := service.ProductService.GetAllProducts()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	var productMaps []map[string]interface{}

	for _, product := range products {
		productMap := map[string]interface{}{
			"id":          product.ID,
			"title":       product.Title,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_Id": product.CategoryID,
			"created_at":  product.CreatedAt,
		}

		productMaps = append(productMaps, productMap)
	}

	context.JSON(http.StatusOK, productMaps)
}

// UpdateProduct godoc
// @Summary Update a Product.
// @Tags Products
// @Accept json
// @Produce json
// @Param productId path int true "Product ID"
// @Param model.Product body model.ProductUpdate true "Product object to be updated"
// @Success 200 {object} model.User "Product updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /products/{productId} [put]
func UpdateProduct(context *gin.Context) {
	var productUpdated model.ProductUpdate

	if err := context.ShouldBindJSON(&productUpdated); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	id, _ := helper.GetIdParam(context, "productId")

	productResponse, err := service.ProductService.UpdateProduct(&productUpdated, id)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result := map[string]interface{}{
		"id": productResponse.ID,
		"title": productResponse.Title,
		"price": productResponse.Price,
		"stock": productResponse.Stock,
		"CategoryId": productResponse.CategoryID,
		"createdAt": productResponse.CreatedAt,
		"updatedAt": productResponse.UpdatedAt,
	}

	context.JSON(http.StatusOK, gin.H{
		"product": result,
	})
}

// DeleteProduct godoc
// @Summary Delete a Product.
// @Tags Products
// @Accept json
// @Produce json
// @Param productId path int true "Product ID"
// @Success 200 {object} model.Product "Product deleted successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /products/{productId} [delete]
func DeleteProduct(context *gin.Context) {
	productId, _ := helper.GetIdParam(context, "productId")

	err := service.ProductService.DeleteProduct(productId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Product has been successfully deleted",
	})
}
