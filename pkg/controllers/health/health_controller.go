package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Shallow(c *gin.Context) {
	c.String(http.StatusOK, "READY \n")
}
