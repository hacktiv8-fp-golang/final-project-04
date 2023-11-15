package middleware

import (
	"final-project-04/internal/database"
	"final-project-04/internal/helper"
	"final-project-04/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		userData := context.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))

		db := database.GetDB()

		var user model.User

		err := db.First(&user, userId).Error

		if err != nil {
			errorHandler := helper.NotFound("Data Not found")
			context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
			return
		}

		if user.Role != "admin" {
			err := helper.Unauthorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
