package services

import (
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if validateError := user.Validate(); validateError != nil {
		return nil,validateError
	}

	return nil, nil

}
