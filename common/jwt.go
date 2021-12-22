package common

import (
	"github.com/dgrijalva/jwt-go"
	"lovelm/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lovelm.gin",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwLCJleHAiOjE2NDA3NzAyNzYsImlhdCI6MTY0MDE2NTQ3NiwiaXNzIjoibG92ZWxtLmdpbiIsInN1YiI6InVzZXIgdG9rZW4ifQ.ELsUL0v2YWg2tKnzoudkXcMsjhOddxSjmTs8PUoVBGc"

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
