package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("secret_key")

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(JwtKey)
}

func ParseToken(tokenStr string) (uint, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	return uint(claims["user_id"].(float64)), nil
}