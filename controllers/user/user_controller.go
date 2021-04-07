package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/regul4rj0hn/bookstore-users-api/domain/users"
	"github.com/regul4rj0hn/bookstore-users-api/services"
)

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}

	result, saveErr := services.Create(user)
	if saveErr != nil {
		fmt.Println(saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}
