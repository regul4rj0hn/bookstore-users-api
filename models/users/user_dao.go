package users

import (
	"github.com/regul4rj0hn/bookstore-users-api/data/psql/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/dates"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
	"github.com/regul4rj0hn/bookstore-users-api/utils/postgres"
)

const (
	querySelectUser = "SELECT * FROM public.user WHERE id = $1;"
	queryInsertUser = "INSERT INTO public.user (first_name, last_name, email, created_on) VALUES ($1, $2, $3, $4) RETURNING id;"
	queryUpdateUser = "UPDATE public.user SET first_name=$1, last_name=$2, email=$3 WHERE id=$4;"
	queryDeleteUser = "DELETE FROM public.user WHERE id=$1;"
)

func (user *User) Get() *errors.Response {
	stmt, err := users.DB.Prepare(querySelectUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedOn); err != nil {
		return postgres.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.Response {
	user.CreatedOn = dates.GetNowString()
	err := users.DB.QueryRow(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreatedOn).Scan(&user.Id)
	if err != nil {
		return postgres.ParseError(err)
	}

	return nil
}

func (user *User) Edit() *errors.Response {
	stmt, err := users.DB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return postgres.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.Response {
	stmt, err := users.DB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return postgres.ParseError(err)
	}

	return nil
}
