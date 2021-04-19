package users

import (
	"fmt"
	"strings"

	"github.com/regul4rj0hn/bookstore-users-api/data/psql/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/dates"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

const (
	errorEmptyResult = "no rows in result set"
	errorUniqueEmail = "user_email_key"

	queryInsertUser = "INSERT INTO public.user (first_name, last_name, email, created_on) VALUES ($1, $2, $3, $4) RETURNING id;"
	queryGetUser    = "SELECT * FROM public.user WHERE id = $1"
)

func (user *User) Get() *errors.Response {
	stmt, err := users.DB.Prepare(queryGetUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedOn); err != nil {
		if strings.Contains(err.Error(), errorEmptyResult) {
			return errors.NotFound(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.InternalServerError(err.Error())
	}

	return nil
}

func (user *User) Save() *errors.Response {
	user.CreatedOn = dates.GetNowString()
	err := users.DB.QueryRow(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreatedOn).Scan(&user.Id)
	if err != nil {
		if strings.Contains(err.Error(), errorUniqueEmail) {
			return errors.Conflict(fmt.Sprintf("email %s is already registered", user.Email))
		}
		return errors.InternalServerError(err.Error())
	}

	return nil
}
