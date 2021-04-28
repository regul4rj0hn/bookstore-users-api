package app

import (
	"github.com/gin-gonic/gin"
	"github.com/regul4rj0hn/bookstore-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("users service API starting...")
	router.Run()
}
