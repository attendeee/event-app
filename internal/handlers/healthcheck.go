package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary      Basic Ping
// @Description  Basic Ping
// @Tags         healthcheck
// @Accept       json
// @Produce      json
// @Router       /ping [get]
func Ping(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
