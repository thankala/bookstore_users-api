package error_parsing_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/thankala/bookstore_users-api/logger"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"strings"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "record not found") {
			logger.Error("User error", err)
			return errors.NewNotFoundError("No record matching given ID")
		}
		logger.Error("Error parsing database response", err)
		return errors.NewInternalError("Error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		logger.Error(fmt.Sprintf("MySQL Error %d:", sqlErr.Number), sqlErr)
		return errors.NewBadRequestError("Email already in use")
	}
	logger.Error("Internal Database Error", err)
	return errors.NewInternalError("Error processing request")
}
