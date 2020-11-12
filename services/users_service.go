package services

import (
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if validateError := user.Validate(); validateError != nil {
		return nil, validateError
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := users.User{Id: userId}
	if err :=user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
