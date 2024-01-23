package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("secret_key")

type JWTClaim struct {
	ID			uint
	Username 	string
	Email 		string
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, username string, email string) (token string, expTime time.Time,err error){
	expTime = time.Now().Add(time.Hour * 2)
	claims := JWTClaim{
		ID: id,
		Username: username,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "task-5_restapi_golang-project",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// MENDEKLARASIKAN ALGORITMA YANG DIGUNAKAN UNTUK SIGNING
	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SIGNED TOKEN
	token, err = tokenAlgorithm.SignedString(JWT_KEY)
	return
}

func ParseToken(tokenString string) (claims *JWTClaim, token *jwt.Token, err error) {
	token, err = jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil {
		return
	}

	claims, isParsed := token.Claims.(*JWTClaim)
	if !isParsed {
		err = errors.New("couldn't parse claims")
		return
	}

	return
}