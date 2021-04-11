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
	CreatedOn string `json:"created_on"`
}

func (user *User) Validate() *errors.Response {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}
	return nil
}
