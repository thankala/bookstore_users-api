package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"strings"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "record not found") {
			return errors.NewNotFoundError("No record matching given ID")
		}
		return errors.NewInternalError("Error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Email already in use")
	}
	return errors.NewInternalError("Error processing request")
}
