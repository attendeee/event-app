package routes

import (
	handler "github.com/attendeee/event-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {
	// Define a simple GET endpoint
	r.GET("/ping", handler.Ping)

	r.GET("/user/all", handler.GetAllUsers)
	r.POST("/user/register", handler.RegisterUser)
	r.PUT("/user/info", handler.UpdateUserInfo)
	r.PUT("/user/password", handler.UpdateUserPassword)
	r.DELETE("/user/:userId", handler.DeleteUserById)
	r.POST("/user/auth", handler.AuthUser)

	r.GET("/event/owner/:ownerId", handler.GetEventsByOwner)
	r.GET("/event/find/:name", handler.GetEventsByName)

	r.POST("/event", handler.CreateEvent)
	r.PUT("/event", handler.UpdateEventInfo)
	r.DELETE("/event/:eventId", handler.DeleteEventByid)

	r.GET("/event/:eventId/attendees", handler.GetAttendees)
	r.POST("/event/addAttendee", handler.AddAttendee)
	r.DELETE("/event/:eventId/attendee/:attendeeId", handler.DeleteAttendee)

}
