package postgres

import (
	"strings"

	"github.com/lib/pq"
	"github.com/regul4rj0hn/bookstore-users-api/pkg/utils/errors"
)

const (
	ErrorEmptyResultSet = "no rows in result set"
)

func ParseError(err error) *errors.Response {
	sqlErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), ErrorEmptyResultSet) {
			return errors.NotFound("no record matching that ID")
		}
		return errors.InternalServerError("error parsing database server response")
	}

	switch sqlErr.Code {
	case "23505":
		return errors.Conflict("duplicate key")
	}

	return errors.InternalServerError("error when processing the request")
}
