package jwtTool_test

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/robbailey3/jwt-tool/jwtTool"
)

func TestCreateToken(t *testing.T) {
	token, err := jwtTool.CreateToken("user1", "secret")

	if err != nil {
		t.Errorf("An error happened: %s", err.Error())
	}

	jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(t)
		return "", nil
	})
}
