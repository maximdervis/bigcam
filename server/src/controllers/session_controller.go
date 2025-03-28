package controllers

import (
	"context"
	"net/http"
	"server/src/db"
	"server/src/util"

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
	userId, exists := ctx.Get("userID")
	if !exists {
		util.SetInternalErrorStatus(ctx, "Failed to load user_id, not authorized?")
	}
	type Request struct {
		GymId    int64 `json:"gym_id" binding:"required"`
		CameraId int64 `json:"camera_id" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}
	err := cc.db.InsertSession(ctx, db.InsertSessionParams{
		UserID:   userId.(int64),
		GymID:    payload.GymId,
		CameraID: payload.CameraId,
	})
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *SessionController) FinishSession(ctx *gin.Context) {
	// TODO: Проверять что корректный пользователь
	type Request struct {
		SessionId int64 `json:"session_id" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}
	cc.db.CloseSession(ctx, payload.SessionId)
	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *SessionController) ListSessions(ctx *gin.Context) {
	userId, exists := ctx.Get("userID")
	if !exists {
		util.SetInternalErrorStatus(ctx, "Failed to load user_id, not authorized?")
	}
	openedSessions, err := cc.db.SelectOpenedSessions(ctx, userId.(int64))
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}
	var response []gin.H
	for _, item := range openedSessions {
		response = append(response, gin.H{
			"session_id": item.ID,
			"camera_id":  item.CameraID,
			"gym_id":     item.GymID,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"sessions": response,
	})
}
