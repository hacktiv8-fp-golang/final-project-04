package helper

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "seed"

func GenerateToken(id int, email string, role string) (string, Error) {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
		"role": role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))

	if err != nil {
		return "", InternalServerError("Failed to generate token")
	}

	return signedToken, nil
}

func VerifyToken(context *gin.Context) (interface{}, Error) {
	headerToken := context.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, Unauthorized("You are not logged in. Please log in to access this resource.")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("You are not logged in. Please log in to access this resource.")
		}

		return []byte(secretKey), nil
	})

	verifiedToken, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, Unauthorized("You are not logged in. Please log in to access this resource.")
	}

	return verifiedToken, nil
}
