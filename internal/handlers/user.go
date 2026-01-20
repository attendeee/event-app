package handler

import (
	"context"
	"net/http"
	"time"

	database "github.com/attendeee/event-app/internal/database/compiled-sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	u := &database.CreateUserParams{}

	err := c.BindJSON(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	/* Todo: make something with this line */
	q := database.Queries{}

	u.Password = string(hash)

	_, err = q.CreateUser(context.Background() /* Todo: make something with this line */, *u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})

}

func AuthUser(c *gin.Context) {
	u := &database.CreateUserParams{}

	err := c.BindJSON(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	/* Todo: make something with this line */
	q := database.Queries{}

	existingUser, err := q.GetUserByEmail(context.Background() /* Todo: make something with this line */, u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(u.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": existingUser.ID,
		"expr":   time.Now().Add(72 * time.Hour), /* Todo: make this line configurable */
	})

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func GetAllUsers(c *gin.Context) {
	/* Todo: make something with this line */
	q := database.Queries{}

	users, err := q.GetAllUsers(context.Background() /* Todo: make something with this line */)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusFound, gin.H{"users": users})

}
