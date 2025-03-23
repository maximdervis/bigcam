package controllers

import (
	"context"
	"net/http"
	"server/src/db"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	db  *db.Queries
	ctx context.Context
}

func NewSessionController(db *db.Queries, ctx context.Context) *SessionController {
	return &SessionController{db, ctx}
}

func (cc *SessionController) CreateSession(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *SessionController) FinishSession(ctx *gin.Context) {
	type Request struct {
		SessionId string `json:"session_id" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *SessionController) ListSessions(ctx *gin.Context) {
	type Request struct {
		SessionId string `json:"session_id" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
