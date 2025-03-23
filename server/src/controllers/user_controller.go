package controllers

import (
	"context"
	"net/http"
	"server/src/db"

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

	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *UserController) VerifySignUpCode(ctx *gin.Context) {
	type Request struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
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

	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *UserController) RefreshAuthToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"email": "maximdervis@yandex-team.ru",
		"name":  "Dervis Maksim Vladimirovich",
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
