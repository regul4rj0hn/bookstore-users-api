package users

import (
	"strings"

	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Status    string `json:"status"`
	CreatedOn string `json:"created_on"`
}

func (user *User) Validate() *errors.Response {
	user.CreatedOn = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}
	return nil
}
