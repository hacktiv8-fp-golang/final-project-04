package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserRegister godoc
// @Summary Register a new user
// @Description Endpoint to register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param model.UserRegister body model.UserRegister true "create user"
// @Success 201 {object} model.UserRegisterResponse
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 422 {object} helper.ErrorResponse "Invalid Request"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /users/register [post]
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

// UserLogin godoc
// @Summary Authenticate a user
// @Description Endpoint to register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param model.LoginCredential body model.LoginCredential true "authenticate user"
// @Success 200 {string} string "OK"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 422 {object} helper.ErrorResponse "Invalid Request"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /users/login [post]
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

// UserTopup godoc
// @Summary Topup user balance
// @Description Endpoint to top-up user balance
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param model.User body model.BalanceUpdate true "Amount to top-up"
// @Success 200 {string} string "OK"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 422 {object} helper.ErrorResponse "Invalid Request"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /users/topup [patch]
func UpdateBalance(context *gin.Context) {
	var balance model.BalanceUpdate

	if err := context.ShouldBindJSON(&balance); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	updatedBalance, err := service.UserService.UpdateBalance(&balance, userId)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp %d", updatedBalance),
	})
}
