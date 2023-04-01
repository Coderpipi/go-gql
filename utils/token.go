package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go-gql/model"
)

const sign = "go-gql-231kvjla3pojflsd"

type UserClaims struct {
	Name  string `json:"name"`
	Type  int8   `json:"type"`
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.User) (string, error) {
	userClaims := &UserClaims{
		Name:             user.UserName,
		Type:             user.Type,
		Phone:            user.Phone,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString([]byte(sign))
	if err != nil {
		panic("Generate Token Failed")
		return "", err
	}
	return tokenString, nil

}
