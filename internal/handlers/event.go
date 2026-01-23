package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	database "github.com/attendeee/event-app/internal/database/compiled-sql"
	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {

	e := database.CreateEventParams{}

	err := c.BindJSON(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	q := database.Queries{}

	err = q.CreateEvent(context.Background(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"event": e})

}

func GetEventsByOwner(c *gin.Context) {

	p := c.Param("ownerId")

	pInt, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id := sql.NullInt64{Int64: int64(pInt), Valid: true}

	q := database.Queries{}

	foundEvents, err := q.GetEventByOwner(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(foundEvents) < 1 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusFound, gin.H{"events": foundEvents})

}

func GetEventsByName(c *gin.Context) {

	name := c.Param("name")

	q := database.Queries{}

	foundEvents, err := q.GetEventByName(context.Background(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(foundEvents) < 1 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusFound, gin.H{"events": foundEvents})

}

func UpdateEventInfo(c *gin.Context) {
	up := database.UpdateEventParams{}
	err := c.BindJSON(up)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	/* Todo: make something with this line */
	q := database.Queries{}

	err = q.UpdateEvent(context.Background() /* Todo: make something with this line */, up)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}

func DeleteEventByid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	/* Todo: make something with this line */
	q := database.Queries{}

	err = q.DeleteEventById(context.Background() /* Todo: make something with this line */, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}
