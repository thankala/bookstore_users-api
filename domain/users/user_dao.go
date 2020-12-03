package users

import (
	"fmt"
	"github.com/thankala/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/thankala/bookstore_users-api/utils/error_parsing_utils"
	"github.com/thankala/bookstore_users-api/utils/errors"
)

func (user *User) Save() *errors.RestError {
	result := bookstore_users.Client.Create(&user)
	if result.Error != nil {
		return error_parsing_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Get() *errors.RestError {
	result := bookstore_users.Client.First(&user, user.ID)
	if result.Error != nil {
		return error_parsing_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Update() *errors.RestError {
	result := bookstore_users.Client.Save(&user)
	if result.Error != nil {

		return error_parsing_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	result := bookstore_users.Client.Delete(&user)
	if result.Error != nil {
		return error_parsing_utils.ParseError(result.Error)
	}
	return nil
}

func (users *Users) FindByStatus(status string) *errors.RestError {
	result := bookstore_users.Client.Where("status = ?", status).Find(&users)
	if result.Error != nil {
		return error_parsing_utils.ParseError(result.Error)
	}
	if len(*users) == 0 {
		return errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return nil
}

func (user *User) FindByEmailAndPassword() *errors.RestError {
	result := bookstore_users.Client.Where(&User{Email: user.Email, Password: user.Password}).First(&user)
	if result.Error != nil {
		return error_parsing_utils.ParseError(result.Error)
	}
	return nil
}
