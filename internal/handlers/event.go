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

// @Summary      Create event
// @Description  Create event
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event [post]
func CreateEvent(c *gin.Context) {

	e := database.CreateEventParams{}

	err := c.BindJSON(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = dbConn.Query.CreateEvent(dbConn.Context, e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"event": e})

}

// @BasePath /api/v1

// @Summary      Get events by owner id
// @Description  Get events by owner id
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/{ownerId} [get]
func GetEventsByOwner(c *gin.Context) {

	p := c.Param("ownerId")

	pInt, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id := sql.NullInt64{Int64: int64(pInt), Valid: true}

	foundEvents, err := dbConn.Query.GetEventByOwner(dbConn.Context, id)
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

// @BasePath /api/v1

// @Summary      Get events by name
// @Description  Get events by name
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/{name} [get]
func GetEventsByName(c *gin.Context) {

	name := c.Param("name")

	foundEvents, err := dbConn.Query.GetEventByName(dbConn.Context, name)
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

// @BasePath /api/v1

// @Summary      Update event info
// @Description  Update event info
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/ [put]
func UpdateEventInfo(c *gin.Context) {
	up := database.UpdateEventParams{}
	err := c.BindJSON(up)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = dbConn.Query.UpdateEvent(dbConn.Context, up)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}

// @BasePath /api/v1

// @Summary      Delete event
// @Description  Delete event
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /event/{eventId} [delete]
func DeleteEventByid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = dbConn.Query.DeleteEventById(dbConn.Context, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}
