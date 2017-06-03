package service

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func GetUidByToken(jwt_string string)  string{
	token, err := jwt.Parse(jwt_string, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123"), nil
	})

	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		uid := fmt.Sprintf("%v",claims["uid"])
		return uid
	}
	return ""

}
