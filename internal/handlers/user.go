package handler

import (
	"net/http"
	"strconv"
	"time"

	database "github.com/attendeee/event-app/internal/database/compiled-sql"
	dbConn "github.com/attendeee/event-app/internal/database/conn"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Summary      Register user
// @Description  Register user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 request body database.CreateUserParams true "User information"
// @Router       /user/register [post]
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

	u.Password = string(hash)

	_, err = dbConn.Query.CreateUser(dbConn.Context, *u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})

}

// @BasePath /api/v1

// @Summary      Authorize user
// @Description  Authorize user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 request body database.CreateUserParams true "User information"
// @Router       /user/auth [post]
func AuthUser(c *gin.Context) {
	u := &database.CreateUserParams{}

	err := c.BindJSON(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	existingUser, err := dbConn.Query.GetUserByEmail(dbConn.Context, u.Email)
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

// @BasePath /api/v1

// @Summary      Get all users
// @Description  Get all users
// @Tags         user
// @Produce      json
// @Router       /user/all [get]
func GetAllUsers(c *gin.Context) {

	users, err := dbConn.Query.GetAllUsers(dbConn.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(users) < 1 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusFound, gin.H{"users": users})

}

// @BasePath /api/v1

// @Summary      Update user info
// @Description  Update user info
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 request body database.UpdateUserInfoParams true "User information to update"
// @Router       /user/info [put]
func UpdateUserInfo(c *gin.Context) {
	up := database.UpdateUserInfoParams{}
	err := c.BindJSON(up)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = dbConn.Query.UpdateUserInfo(dbConn.Context, up)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}

// @BasePath /api/v1

// @Summary      Update user password
// @Description  Update user password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 request body database.UpdateUserPasswordParams true "User password to update"
// @Router       /user/password [put]
func UpdateUserPassword(c *gin.Context) {
	up := database.UpdateUserPasswordParams{}
	err := c.BindJSON(up)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = dbConn.Query.UpdateUserPassword(dbConn.Context, up)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}

// @BasePath /api/v1

// @Summary      Delete user
// @Description  Delete user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /user/{id} [delete]
func DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = dbConn.Query.DeleteUserById(dbConn.Context, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)

}
