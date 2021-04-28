package app

import (
	"github.com/regul4rj0hn/bookstore-users-api/pkg/controllers/health"
	"github.com/regul4rj0hn/bookstore-users-api/pkg/controllers/user"
)

func mapUrls() {
	router.GET("/health", health.Shallow)

	router.POST("/user", user.Create)
	router.GET("/user/:id", user.Get)
	router.PUT("/user/:id", user.Update)
	router.PATCH("/user/:id", user.Update)
	router.DELETE("/user/:id", user.Delete)

	router.GET("/internal/user/search", user.Search)
}
