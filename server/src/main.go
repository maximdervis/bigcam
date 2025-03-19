package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"server/src/controllers"
	"server/src/routes"
	"server/src/util"

	dbCon "server/src/db"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	ContactController controllers.GymController
	ContactRoutes     routes.Routes
)


func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Doing well, by now"})
	})

	ContactRoutes.Route(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))
}

