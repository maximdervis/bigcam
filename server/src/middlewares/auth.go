package middlewares

import (
	"net/http"
	"server/src/models"
	"server/src/util"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var accessSecretKey = []byte("my_secret_key_access")
var refreshSecretKey = []byte("my_secret_key_refresh")

func GetAccessSignedToken(userId int64, email string) (string, error) {
	expiresAt := time.Now().Add(5 * time.Minute)
	token, err := getSignedToken(userId, email, expiresAt, accessSecretKey)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func GetRefreshSignedToken(userId int64, email string) (string, error) {
	expiresAt := time.Now().Add(48 * time.Hour)
	token, err := getSignedToken(userId, email, expiresAt, refreshSecretKey)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func getSignedToken(userId int64, email string, exiresAt time.Time, secret []byte) (string, error) {
	token := getToken(userId, email, exiresAt)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func getToken(userId int64, email string, expiresAt time.Time) (*jwt.Token) {
	return jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			models.Claims{
				UserId: userId,
				StandardClaims: jwt.StandardClaims{
				    Subject:   email,
				    ExpiresAt: expiresAt.Unix(),
				},
			},
		)
}

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusForbidden, gin.H{"code": "ACCESS_DENIED", "message": "Missing Authorization header"})
			c.Abort()
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			c.JSON(http.StatusForbidden, gin.H{"code": "ACCESS_DENIED", "message": "Invalid access token format"})
			c.Abort()
			return
		}

		accessToken := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := util.ParseToken(accessToken, accessSecretKey)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"code": "ACCESS_DENIED", "message": "Failed to parse access token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserId)
		c.Next()
	}
}
