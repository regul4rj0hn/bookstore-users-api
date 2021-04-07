package services

import "github.com/regul4rj0hn/bookstore-users-api/domain/users"

func Create(user users.User) (*users.User, error) {
	return &user, nil
}
