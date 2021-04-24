package services

import (
	"github.com/regul4rj0hn/bookstore-users-api/models/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

func Create(user users.User) (*users.User, *errors.Response) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func Get(userId int64) (*users.User, *errors.Response) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func Update(user users.User) (*users.User, *errors.Response) {
	current, err := Get(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email
	if err := current.Edit(); err != nil {
		return nil, err
	}
	return current, nil
}
