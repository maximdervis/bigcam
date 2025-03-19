package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"server/src/db"
	"server/src/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GymController struct {
	db  *db.Queries
	ctx context.Context
}

func NewGymController(db *db.Queries, ctx context.Context) *GymController {
	return &GymController{db, ctx}
}

func (cc *GymController) CreateGym(ctx *gin.Context) {
	var payload *schemas.CreateGym

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	err := cc.db.InsertGym(ctx, payload.Name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully created contact"})
}

func (cc *GymController) GetGym(ctx *gin.Context) {
	gymIdStr := ctx.Param("gym_id")
	gymIdInt, err := strconv.Atoi(gymIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed to parse integer from gym_id", "error": err.Error()})
	}
	gym, err := cc.db.SelectGymInfo(ctx, int32(gymIdInt));
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve gym with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving gym", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrived id", "gym": gym})
}
