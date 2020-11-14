package users

import (
	"fmt"
	"github.com/thankala/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/thankala/bookstore_users-api/utils/date_utils"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"strings"
)

func (user *User) Get() *errors.RestError {
	result := bookstore_users.Client.First(&user, user.ID)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "record not found") {
			return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
		}
		return errors.NewInternalError(fmt.Sprintf("Error when trying to get user %d: %s", user.ID, result.Error.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	result := bookstore_users.Client.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Error 1062") {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewInternalError(fmt.Sprintf("Error when trying to save user: %s", result.Error.Error()))

	}
	bookstore_users.Client.Model(&user).Update("UpdatedAt", date_utils.GetNow())
	bookstore_users.Client.Model(&user).Update("CreatedAt", date_utils.GetNow())
	return nil
}

func (user *User) Delete() *errors.RestError {
	bookstore_users.Client.Model(&user).Update("DeletedAt", date_utils.GetNow())
	return nil
}
