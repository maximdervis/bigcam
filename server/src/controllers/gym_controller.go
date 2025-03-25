package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"server/src/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type GymController struct {
	db    *db.Queries
	redis *redis.Client
	ctx   context.Context
}

func NewGymController(db *db.Queries, redis *redis.Client, ctx context.Context) *GymController {
	return &GymController{db, redis, ctx}
}

func (cc *GymController) CreateGym(ctx *gin.Context) {
	type CreateGym struct {
		Name string `json:"name" binding:"required"`
	}
	var payload *CreateGym

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "error": err.Error()})
		return
	}

	uuid := uuid.New()
	insert_params := db.InsertGymParams{
		Name:    payload.Name,
		AuthKey: uuid.String(),
	}
	err := cc.db.InsertGym(ctx, insert_params)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"code": "INTERNAL_ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"auth_key": insert_params.AuthKey})
}

func (cc *GymController) GetGym(ctx *gin.Context) {
	gymIdStr := ctx.Param("gym_id")
	gymIdInt, err := strconv.Atoi(gymIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": err.Error()})
	}
	gym, err := cc.db.SelectGymInfo(ctx, int64(gymIdInt))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"code": "NOT_FOUND", "message": "Failed to retrieve gym with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"code": "INTERNAL_ERROR", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"gym": gym})
}

func (cc *GymController) LocalGymAssign(ctx *gin.Context) {
	type LocalGymAssign struct {
		AuthKey string `json:"auth_key" binding:"required"`
	}
	var payload *LocalGymAssign

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "message": err.Error()})
		return
	}

	gym_id, err := cc.db.SelectGymIdByAuthKey(ctx, payload.AuthKey)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"code": "INTERNAL_ERROR", "message": err.Error()})
		return
	}

	client_ip := ctx.ClientIP()
	fmt.Println("Got client_ip:", client_ip)
	cc.redis.Set(ctx, strconv.FormatInt(gym_id, 10), client_ip, 0)

	ctx.JSON(http.StatusOK, gin.H{})
}
