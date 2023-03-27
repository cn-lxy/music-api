package tools

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

// This package have function tools of generate jwt token, verify jwt token.

// GenerateToken generate jwt token
func GenerateToken(id uint64) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.SigningMethodHS256)
	// set claims
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = id
	// generate encoded token and send it as response.
	token, err := t.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyToken verify jwt token
func VerifyToken(tokenString string) (uint64, error) {
	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	// validate token
	if !token.Valid {
		return 0, err
	}
	// get claims
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(uint64)
	return id, nil
}
