package users

import (
	"fmt"

	"github.com/regul4rj0hn/bookstore-users-api/data/psql/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
	"github.com/regul4rj0hn/bookstore-users-api/utils/postgres"
)

const (
	querySelectUser       = "SELECT first_name, last_name, email, status, created_on FROM public.user WHERE id = $1;"
	queryInsertUser       = "INSERT INTO public.user (first_name, last_name, email, password, status, created_on) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;"
	queryUpdateUser       = "UPDATE public.user SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"
	queryDeleteUser       = "DELETE FROM public.user WHERE id = $1;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, status, created_on FROM public.user WHERE status = $1;"
)

func (user *User) Get() *errors.Response {
	stmt, err := users.DB.Prepare(querySelectUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedOn); err != nil {
		return postgres.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.Response {
	err := users.DB.QueryRow(queryInsertUser, user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedOn).Scan(&user.Id)
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

func (user *User) FindByStatus(status string) ([]User, *errors.Response) {
	stmt, err := users.DB.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedOn); err != nil {
			return nil, postgres.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFound(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
