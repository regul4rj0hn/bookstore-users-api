package users

import (
	"strings"

	"github.com/regul4rj0hn/bookstore-users-api/pkg/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	CreatedOn string `json:"created_on"`
}

type Users []User

func (user *User) Validate() *errors.Response {
	user.CreatedOn = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.BadRequest("invalid password")
	}
	return nil
}
