package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/cn-lxy/music-api/models"
	"github.com/golang-jwt/jwt/v4"
)

// This package have function tools of generate jwt token, verify jwt token.

// GenerateToken generate jwt token
func GenerateToken(u *models.User) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":       u.Id,
		"nickname": u.NickName,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 72 hours expiration
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return tokenString, nil
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
