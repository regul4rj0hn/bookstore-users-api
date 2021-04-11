package app

import (
	"github.com/regul4rj0hn/bookstore-users-api/controllers/health"
	"github.com/regul4rj0hn/bookstore-users-api/controllers/user"
)

func mapUrls() {
	router.GET("/health", health.Shallow)

	router.POST("/user", user.Create)
	router.GET("/user/:id", user.Get)
}
