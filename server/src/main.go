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
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	GymController       controllers.GymController
	GymCameraController controllers.GymCameraController
	UserController      controllers.UserController
	GymRoute            routes.GymRoute
	GymCamaeraRoute     routes.GymCameraRoute
	UserRoute           routes.UserRoute
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

	redis_client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redis_client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("could not connect to redis: %v", err)
	}
	fmt.Println("Redist connected successfully: ", pong)

	GymController = *controllers.NewGymController(db, ctx)
	UserController = *controllers.NewUserController(db, ctx)
	GymCameraController = *controllers.NewGymCameraController(db, ctx)
	GymRoute = routes.NewGymRoute(GymController)
	GymCamaeraRoute = routes.NewGymCameraRoute(GymCameraController)
	UserRoute = routes.NewUserRoute(UserController)

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
	UserRoute.Route(router)
	GymCamaeraRoute.Route(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))
}
