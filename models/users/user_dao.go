package users

import (
	"fmt"
	"strings"

	"github.com/regul4rj0hn/bookstore-users-api/data/psql/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/dates"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

const (
	constraintUniqueEmail = "user_email_key"
	queryInsertUser       = "INSERT INTO public.user (first_name, last_name, email, created_on) VALUES ($1, $2, $3, $4) RETURNING id;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.Response {
	if err := users.DB.Ping(); err != nil {
		panic(err)
	}

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
	user.CreatedOn = dates.GetNowString()
	err := users.DB.QueryRow(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreatedOn).Scan(&user.Id)
	if err != nil {
		if strings.Contains(err.Error(), constraintUniqueEmail) {
			return errors.Conflict(fmt.Sprintf("email %s is already registered", user.Email))
		}
		return errors.InternalServerError(err.Error())
	}

	return nil
}
