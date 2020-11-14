package users

import (
	"fmt"
	"github.com/thankala/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/thankala/bookstore_users-api/utils/errors"
)

func (user *User) Get() *errors.RestError {
	result := bookstore_users.Client.First(&user, user.ID)
	if result.Error != nil {
		return errors.NewInternalError(fmt.Sprintf("Error when trying to get user: %s", result.Error.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	result := bookstore_users.Client.Create(&user)
	if result.Error != nil {
		return errors.NewInternalError(fmt.Sprintf("Error when trying to save user: %s", result.Error.Error()))
	}
	return nil
}
