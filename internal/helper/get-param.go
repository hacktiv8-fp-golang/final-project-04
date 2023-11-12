package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParam(context *gin.Context, paramName string) (int, Error) {
	id, err := strconv.Atoi(context.Param(paramName))

	if err != nil {
		return int(0), BadRequest("Invalid id params")
	}

	return int(id), nil
}
