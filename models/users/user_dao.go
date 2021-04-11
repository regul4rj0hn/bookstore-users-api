package users

import (
	"fmt"

	"github.com/regul4rj0hn/bookstore-users-api/utils/dates"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.Response {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFound(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.CreatedOn = result.CreatedOn

	return nil
}

func (user *User) Save() *errors.Response {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.Conflict(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.Conflict(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.CreatedOn = dates.GetNowString()

	usersDB[user.Id] = user
	return nil
}
