package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	database "github.com/attendeee/event-app/internal/database/compiled-sql"
	dbConn "github.com/attendeee/event-app/internal/database/conn"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary      Get attendees
// @Description  Get attendees for event
// @Tags         attendee
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/{eventId}/attendees [get]
func GetAttendees(c *gin.Context) {

	p := c.Param("eventId")

	pInt, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id := sql.NullInt64{Int64: int64(pInt), Valid: true}

	foundAttendees, err := dbConn.Query.GetAllAttendeesForEvent(dbConn.Context, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(foundAttendees) < 1 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusFound, gin.H{"attendees": foundAttendees})

}

// @BasePath /api/v1

// @Summary      Add attendee
// @Description  Add attendee for event
// @Tags         attendee
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/addAttendee [post]
func AddAttendee(c *gin.Context) {

	a := &database.AddAttendeeParams{}

	err := c.BindJSON(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = dbConn.Query.AddAttendee(dbConn.Context, *a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)

}

// @BasePath /api/v1

// @Summary      Delete attendee
// @Description  Delete attendee from event
// @Tags         attendee
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/{eventId}/attendee/{attendeeId} [delete]
func DeleteAttendee(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	eid, err := strconv.Atoi(c.Param("eventId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	a := database.DeleteAttendeeParams{
		UserID:  sql.NullInt64{Int64: int64(uid), Valid: true},
		EventID: sql.NullInt64{Int64: int64(eid), Valid: true},
	}

	err = dbConn.Query.DeleteAttendee(dbConn.Context, a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}
