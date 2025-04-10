package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var accessSecretKey = []byte("my_secret_key_access")
var refreshSecretKey = []byte("my_secret_key_refresh")

type AuthData struct {
	AccessKey  string
	RefreshKey string
}

func GetSignedTokens(userId string) (*AuthData, error) {
	var err error
	accessKey, err := GetAccessSignedToken(userId)
	if err != nil {
		return nil, nil
	}
	refreshKey, err := GetRefreshSignedToken(userId)
	if err != nil {
		return nil, nil
	}
	authData := &AuthData{
		AccessKey:  accessKey,
		RefreshKey: refreshKey,
	}
	return authData, nil
}

func GetAccessSignedToken(userId string) (string, error) {
	expiresAt := time.Now().Add(100 * time.Minute)
	token, err := getSignedToken(userId, expiresAt, accessSecretKey)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func GetRefreshSignedToken(userId string) (string, error) {
	expiresAt := time.Now().Add(48 * time.Hour)
	token, err := getSignedToken(userId, expiresAt, refreshSecretKey)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseAccessToken(tokenString string) (claims *jwt.StandardClaims, err error) {
	return parseToken(tokenString, accessSecretKey)
}

func ParseRefreshToken(tokenString string) (claims *jwt.StandardClaims, err error) {
	return parseToken(tokenString, refreshSecretKey)
}

func parseToken(tokenString string, sectetKey []byte) (claims *jwt.StandardClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sectetKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

func getSignedToken(userId string, exiresAt time.Time, secret []byte) (string, error) {
	token := getToken(userId, exiresAt)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func getToken(userId string, expiresAt time.Time) *jwt.Token {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Subject:   userId,
			ExpiresAt: expiresAt.Unix(),
		},
	)
}
