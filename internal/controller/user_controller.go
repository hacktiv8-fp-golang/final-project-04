package controller

import (
	"final-project-04/internal/helper"
	"final-project-04/internal/model"
	"final-project-04/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	user.Role = "customer"

	userResponse, err := service.UserService.Register(&user)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id": userResponse.ID,
		"full_name": userResponse.FullName,
		"email": userResponse.Email,
		"password": userResponse.Password,
		"balance": userResponse.Balance,
		"created_at": userResponse.CreatedAt,
	})
}

func Login(context *gin.Context) {
	var userLogin model.LoginCredential

	if err := context.ShouldBindJSON(&userLogin); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	token, err := service.UserService.Login(&userLogin)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
