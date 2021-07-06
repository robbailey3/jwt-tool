package jwtTool

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId, key string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().AddDate(1, 0, 0).Unix(),
	}).SignedString([]byte(key))
}
