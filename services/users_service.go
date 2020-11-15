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

func GetUser(userId uint) (*users.User, *errors.RestError) {
	user := users.User{ID: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" {
		current.LastName = user.LastName
	}
	if user.Email != "" {
		current.Email = user.Email
	}
	if validateError := current.Validate(); validateError != nil {
		return nil, validateError
	}
	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId uint) *errors.RestError {
	user := users.User{ID: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestError) {
	user := users.User{}
	return user.FindByStatus(status)
}
