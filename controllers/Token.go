package controllers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("jwt-key")

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expire := time.Now().Add(1 * time.Hour)

	jwtClaim := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Id:        "1",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	tokenString, err := token.SignedString(JwtKey)

	return tokenString, err
}

func ValidationToken(tokenSign string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenSign, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
	}

	return claims, err
}
