package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"server/src/db"
	"server/src/middlewares"
	"server/src/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserController(db *db.Queries, ctx context.Context) *UserController {
	return &UserController{db, ctx}
}

func (cc *UserController) SignUp(ctx *gin.Context) {
	type Request struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	userAlreadyExists, checkErr := cc.db.ContainsUserWithEmail(ctx, payload.Email)
	if checkErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": checkErr.Error()})
		return
	}

	if userAlreadyExists {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUESTS", "message": "User with the email already exists"})
		return
	}

	user := db.InsertUserInfoParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	var errHash error
	user.Password, errHash = util.GenerateHashPassword(user.Password)
	if errHash != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": errHash.Error()})
		return
	}

	errInsert := cc.db.InsertUserInfo(ctx, user)
	if errInsert != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": errInsert.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *UserController) SignIn(ctx *gin.Context) {
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	userInfo, selectErr := cc.db.SelectUserInfoByEmail(ctx, payload.Email)
	if selectErr != nil {
		if selectErr == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"code": "NOT_FOUND", "message": "Failed to find user with the email"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": selectErr.Error()})
		return
	}

	passwordsMatched := util.CompareHashPassword(payload.Password, userInfo.Password)
	if !passwordsMatched {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": "Invalid password"})
		return
	}

	accessToken, accessErr := middlewares.GetAccessSignedToken(userInfo.ID, userInfo.Email)
	if accessErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": accessErr.Error()})
		return
	}
	refreshToken, refreshErr := middlewares.GetRefreshSignedToken(userInfo.ID, userInfo.Email)
	if refreshErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": refreshErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (cc *UserController) RefreshAuthToken(ctx *gin.Context) {
	type Request struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}
	var refreshSecretKey = []byte("my_secret_key_refresh")
	claims, parseErr := util.ParseToken(payload.RefreshToken, refreshSecretKey)
	if parseErr != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"code": "ACCESS_DENIED", "message": parseErr.Error()})
		return
	}
	accessToken, accessErr := middlewares.GetAccessSignedToken(claims.UserId, claims.Subject)
	if accessErr != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"code": "ACCESS_DENIED", "message": accessErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": payload.RefreshToken,
	})
}

func (cc *UserController) GetUser(ctx *gin.Context) {
	userId, _ := ctx.Get("userID")
	userInfo, _ := cc.db.SelectUserInfo(ctx, userId.(int64))
	ctx.JSON(http.StatusOK, gin.H{
		"email": userInfo.Email,
		"name":  userInfo.Name,
		"Dob":   "2005-05-20",
		"AvaId": "blob:DdksajkdsjaKdsajldksjalKJD",
	})
}

func (cc *UserController) UpdateUser(ctx *gin.Context) {
	type Request struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
		Dob   string `json:"dob,omitempty"`
		AvaId string `json:"ava_id,omitempty"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
