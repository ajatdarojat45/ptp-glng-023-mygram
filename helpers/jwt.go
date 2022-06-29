package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const secretkey = "rahasia"

func GenerateToken(username string) string {
	claims := jwt.MapClaims{
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretkey))

	return signedToken
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	var errResponse = errors.New("Failed to verify token")
	parseToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretkey), nil
	})

	if _, ok := parseToken.Claims.(jwt.MapClaims); !ok && !parseToken.Valid {
		return nil, errResponse
	}
	return parseToken.Claims.(jwt.MapClaims), nil
}
