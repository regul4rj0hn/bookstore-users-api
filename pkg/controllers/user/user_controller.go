package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/regul4rj0hn/bookstore-users-api/pkg/models/users"
	"github.com/regul4rj0hn/bookstore-users-api/pkg/services"
	"github.com/regul4rj0hn/bookstore-users-api/pkg/utils/errors"
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

	result, err := services.UsersService.Create(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.UsersService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
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

	result, updateErr := services.UsersService.Update(user, isPartial)
	if updateErr != nil {
		c.JSON(updateErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if err := services.UsersService.Delete(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
