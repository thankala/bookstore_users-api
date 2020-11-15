package users

import (
	"github.com/thankala/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"github.com/thankala/bookstore_users-api/utils/mysql_utils"
)

func (user *User) Save() *errors.RestError {
	result := bookstore_users.Client.Create(&user)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Get() *errors.RestError {
	result := bookstore_users.Client.First(&user, user.ID)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Update() *errors.RestError {
	result := bookstore_users.Client.Save(&user)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	result := bookstore_users.Client.Delete(&user)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	var users []User
	result := bookstore_users.Client.Where("status = ?", status).Find(&users)
	if result.Error != nil {
		return nil, mysql_utils.ParseError(result.Error)
	}
	return users, nil
}
