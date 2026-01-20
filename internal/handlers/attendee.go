package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	database "github.com/attendeee/event-app/internal/database/compiled-sql"
	"github.com/gin-gonic/gin"
)

func GetAttendees(c *gin.Context) {

	p := c.Param("eventId")

	pInt, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id := sql.NullInt64{Int64: int64(pInt), Valid: true}

	q := database.Queries{}

	foundAttendees, err := q.GetAllAttendeesForEvent(context.Background(), id)
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

func AddAttendee(c *gin.Context) {

	a := &database.AddAttendeeParams{}

	err := c.BindJSON(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	q := database.Queries{}

	err = q.AddAttendee(context.Background(), *a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)

}
