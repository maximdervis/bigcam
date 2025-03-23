package controllers

import (
	"context"
	"net/http"
	"server/src/db"

	"github.com/gin-gonic/gin"
)

type GymCameraController struct {
	db  *db.Queries
	ctx context.Context
}

func NewGymCameraController(db *db.Queries, ctx context.Context) *GymCameraController {
	return &GymCameraController{db, ctx}
}

func (cc *GymCameraController) GetCameras(ctx *gin.Context) {
	type Request struct {
		GymId string `json:"gym_id" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cameras": []gin.H{
			{
				"id":          52,
				"description": "Best cam for watching cum",
				"occupied_by": "maximdervis@yandex-team.ru",
			},
		},
	})
}
