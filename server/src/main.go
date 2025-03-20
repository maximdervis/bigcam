package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"server/src/controllers"
	"server/src/routes"
	"server/src/util"

	dbCon "server/src/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	GymController controllers.GymController
	GymRoute      routes.GymRoute
)

func init() {
	ctx = context.TODO()
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	log.Printf("Connecting using driver %s, source %s", config.DbDriver, config.DbSource)
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	db = dbCon.New(conn)

	fmt.Println("PostgreSql connected successfully...")

	GymController = *controllers.NewGymController(db, ctx)
	GymRoute = routes.NewRoute(GymController)

	server = gin.Default()
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Doing well, by now"})
	})

	GymRoute.Route(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))
}
