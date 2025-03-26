package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"server/src/db"
	"server/src/middlewares"
	"server/src/util"
	"strconv"
	"time"

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
	var err error
	type Request struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var payload *Request
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}

	userAlreadyExists, err := cc.db.ContainsUserWithEmail(ctx, payload.Email)
	if err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}

	if userAlreadyExists {
		util.SetBadRequestStatus(ctx, "User with the email already exists")
		return
	}

	user := db.InsertUserInfoParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	user.Password, err = middlewares.GenerateHashPassword(user.Password)
	if err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}

	err = cc.db.InsertUserInfo(ctx, user)
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *UserController) SignIn(ctx *gin.Context) {
	var err error
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var payload *Request
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	userInfo, err := cc.db.SelectUserInfoByEmail(ctx, payload.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			util.SetNotFoundStatus(ctx, err)
			return
		}
		util.SetInternalErrorStatus(ctx, err)
		return
	}

	passwordsMatched := middlewares.CompareHashPassword(payload.Password, userInfo.Password)
	if !passwordsMatched {
		util.SetBadRequestStatus(ctx, "Invalid password")
		return
	}

	authData, err := middlewares.GetSignedTokens(strconv.FormatInt(userInfo.ID, 10))
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  authData.AccessKey,
		"refresh_token": authData.RefreshKey,
	})
}

func (cc *UserController) RefreshAuthToken(ctx *gin.Context) {
	var err error
	type Request struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
	}
	var payload *Request
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}
	claims, err := middlewares.ParseRefreshToken(payload.RefreshToken)
	if err != nil {
		util.SetAccessDeniedStatusStatus(ctx, err)
		return
	}
	accessToken, err := middlewares.GetAccessSignedToken(claims.Subject)
	if err != nil {
		util.SetAccessDeniedStatusStatus(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": payload.RefreshToken,
	})
}

func (cc *UserController) GetUser(ctx *gin.Context) {
	var err error
	userId, exists := ctx.Get("userID")
	if !exists {
		util.SetInternalErrorStatus(ctx, "Failed to load user_id, not authorized?")
	}
	userInfo, err := cc.db.SelectUserInfo(ctx, userId.(int64))
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"email":     userInfo.Email,
		"name":      userInfo.Name,
		"dob":       userInfo.Dob,
		"avatar_id": userInfo.AvatarID,
	})
}

func (cc *UserController) UpdateUser(ctx *gin.Context) {
	var err error
	userId, exists := ctx.Get("userID")
	if !exists {
		util.SetInternalErrorStatus(ctx, "Failed to load user_id, not authorized?")
	}
	type Request struct {
		Email    *string    `json:"email,omitempty"`
		Name     *string    `json:"name,omitempty"`
		Dob      *time.Time `json:"dob,omitempty"`
		AvatarId *string    `json:"avatar_id,omitempty"`
	}
	var payload *Request
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}

	updateData, err := payload.Dob.MarshalJSON()
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}
	err = cc.db.UpdateUserInfo(ctx, db.UpdateUserInfoParams{
		UpdateData: updateData,
		ID:         userId.(int64),
	})
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
