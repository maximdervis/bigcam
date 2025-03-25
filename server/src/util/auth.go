package util

import (
	"server/src/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func ParseToken(tokenString string, key []byte) (claims *models.Claims, err error) {
    token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
        return key, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*models.Claims)

    if !ok {
        return nil, err
    }

    return claims, nil
}

