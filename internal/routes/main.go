package routes

import (
	handler "github.com/attendeee/event-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {
	// Define a simple GET endpoint
	r.GET("/ping", handler.Ping)

	r.GET("/user/all", handler.GetAllUsers)
	r.POST("/user/regisger", handler.RegisterUser)
	r.POST("/user/auth", handler.AuthUser)

	r.GET("/event/owner/:ownerId", handler.GetEventsByOwner)
	r.GET("/event/find/:name", handler.GetEventsByName)

	r.POST("/event", handler.CreateEvent)

	r.GET("/event/:eventId/attendees", handler.GetAttendees)
	r.POST("/event/addAttendee", handler.AddAttendee)

}
