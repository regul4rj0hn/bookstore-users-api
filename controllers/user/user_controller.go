package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/regul4rj0hn/bookstore-users-api/models/users"
	"github.com/regul4rj0hn/bookstore-users-api/services"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

func getUserId(id string) (int64, *errors.Response) {
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, errors.BadRequest("invalid user id")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response := errors.BadRequest(err.Error())
		c.JSON(response.Status, response)
		return
	}

	result, err := services.Create(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response := errors.BadRequest(err.Error())
		c.JSON(response.Status, response)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.Update(user, isPartial)
	if updateErr != nil {
		c.JSON(updateErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
