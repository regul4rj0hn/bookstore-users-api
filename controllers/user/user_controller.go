package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/regul4rj0hn/bookstore-users-api/domain/users"
	"github.com/regul4rj0hn/bookstore-users-api/services"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

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
	c.String(http.StatusNotImplemented, "Not implemented")
}
