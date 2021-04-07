package services

import (
	"github.com/regul4rj0hn/bookstore-users-api/domain/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

func Create(user users.User) (*users.User, *errors.Response) {
	return &user, nil
}
