package main

import (
	"fmt"
	"log"
	"os"

	"serviceTwo/config"
	"serviceTwo/internal/stats"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("CLICKHOUSE_HOST")
	port := os.Getenv("CLICKHOUSE_PORT")
	user := os.Getenv("CLICKHOUSE_USER")
	password := os.Getenv("CLICKHOUSE_PASSWORD")
	database := os.Getenv("CLICKHOUSE_DATABASE")

	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s", user, password, host, port, database)

	Config := config.NewConfig(dsn)

	router := gin.Default()

	stats.NewStatsHandler(router, stats.HandlerDeps{Config: Config})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}
