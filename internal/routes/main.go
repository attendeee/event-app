package routes

import (
	handler "github.com/attendeee/event-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {
	// Define a simple GET endpoint
	r.GET("/ping", handler.Ping)

}
