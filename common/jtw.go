package common

import (
	"gin_vue_practice/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 发放Token
func RleaseToken(user model.User) (string, error){
	expirationTime := time.Now().Add(1 * 24 * time.Hour).Unix()
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt: time.Now().Unix(),
			Issuer: "wl",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		return "", err
	}else {
		return tokenString,err
	}
}

// 解析token

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims :=  &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	return token, claims, err
}
