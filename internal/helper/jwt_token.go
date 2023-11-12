package helper

import "github.com/dgrijalva/jwt-go"

var secretKey = "seed"

func GenerateToken(id int, email string) (string, Error) {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))

	if err != nil {
		return "", InternalServerError("Failed to generate token")
	}

	return signedToken, nil
}
