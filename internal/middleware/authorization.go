package middleware

import (
	"github.com/hacktiv8-fp-golang/final-project-04/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-04/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		userData := context.MustGet("userData").(jwt.MapClaims)
		userRole := userData["role"].(string)

		if userRole != "admin" {
			err := helper.Unauthorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func CategoryAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		categoryId, err := helper.GetIdParam(context, "categoryId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		db := database.GetDB()
		category := model.Category{}

		errMsg := db.First(&category, categoryId).Error
		if errMsg != nil {
			err := helper.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		productId, err := helper.GetIdParam(context, "productId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		db := database.GetDB()
		product := model.Product{}

		errMsg := db.First(&product, productId).Error
		if errMsg != nil {
			err := helper.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
