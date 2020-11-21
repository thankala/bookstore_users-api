package users

import (
	"fmt"
	"github.com/thankala/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/thankala/bookstore_users-api/logger"
	"github.com/thankala/bookstore_users-api/utils/errors"
)

func (user *User) Save() *errors.RestError {
	result := bookstore_users.Client.Create(&user)
	if result.Error != nil {
		logger.Error("Error in dao save user", result.Error)
		return errors.NewInternalError("Database Error")
	}
	return nil
}

func (user *User) Get() *errors.RestError {
	result := bookstore_users.Client.First(&user, user.ID)
	if result.Error != nil {
		logger.Error("Error in dao get user", result.Error)
		return errors.NewInternalError("Database Error")
	}
	return nil
}

func (user *User) Update() *errors.RestError {
	result := bookstore_users.Client.Save(&user)
	if result.Error != nil {
		logger.Error("Error in dao update user", result.Error)
		return errors.NewInternalError("Database Error")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	result := bookstore_users.Client.Delete(&user)
	if result.Error != nil {
		logger.Error("Error in dao delete user", result.Error)
		return errors.NewInternalError("Database Error")
	}
	return nil
}

func (users *Users) FindByStatus(status string) *errors.RestError {
	result := bookstore_users.Client.Where("status = ?", status).Find(&users)
	if result.Error != nil {
		logger.Error("Error in dao findbystatus user", result.Error)
		return errors.NewInternalError("Database Error")
	}
	if len(*users) == 0 {
		return errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return nil
}
