package main

import (
	"database/sql"
	"log"
	"net/http"

	dbCon "gen_server/src/db"
	"gen_server/src/generated"
	"gen_server/src/middlewares"
	"gen_server/src/services"
	"gen_server/src/utils"
	"gen_server/src/views"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func PgConnect(config utils.Config) (queries *dbCon.Queries, conn *sql.DB) {
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Panicf("Could not connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Panicf("Could not ping database: %v", err)
	}

	queries = dbCon.New(conn)
	return queries, conn
}

func RedisConnect(config utils.Config) (r *redis.Client) {
	r = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := r.Ping().Result()
	if err != nil {
		log.Panicf("Could not ping redis: %v", err)
	}

	return r
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not loadconfig: %v", err)
	}
	queries, db := PgConnect(config)
	redis := RedisConnect(config)
	handlers := views.NewHandlers(
		services.NewUserService(db, queries),
		services.NewGymService(queries),
		services.NewLocalGymService(queries, redis),
		services.NewSessionService(queries),
		services.NewCameraService(queries, redis),
	)
	srv, err := api.NewServer(handlers)
	wrappedHandler := middlewares.LoggingMiddleware(middlewares.AuthMiddleware(srv))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(wrappedHandler)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
