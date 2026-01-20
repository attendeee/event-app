package main

import (
	"log"
	"time"

	"github.com/attendeee/event-app/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"http://example.com"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))
	v1 := r.Group("/api/v1")

	routes.V1(v1)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
