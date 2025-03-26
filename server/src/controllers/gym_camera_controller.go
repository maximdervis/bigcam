package controllers

import (
	"context"
	"fmt"
	"net/http"
	"server/src/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type GymCameraController struct {
	db    *db.Queries
	redis *redis.Client
	ctx   context.Context
}

func NewGymCameraController(db *db.Queries, redis *redis.Client, ctx context.Context) *GymCameraController {
	return &GymCameraController{db, redis, ctx}
}

func (cc *GymCameraController) GetCameras(ctx *gin.Context) {
	// TODO: Доступы на получение камер в зале
	type Request struct {
		GymId string `json:"gym_id" binding:"required"`
	}
	var payload *Request
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": err.Error()})
		return
	}

	ip_addr := cc.redis.Get(ctx, payload.GymId)
	fmt.Println("Got client ip address", ip_addr.Val())

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
