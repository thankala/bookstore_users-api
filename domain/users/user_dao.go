package users

import (
	"github.com/thankala/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/thankala/bookstore_users-api/utils/date_utils"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"github.com/thankala/bookstore_users-api/utils/mysql_utils"
)

func (user *User) Get() *errors.RestError {
	result := bookstore_users.Client.First(&user, user.ID)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	result := bookstore_users.Client.Create(&user)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	bookstore_users.Client.Model(&user).Update("UpdatedAt", date_utils.GetNow())
	bookstore_users.Client.Model(&user).Update("CreatedAt", date_utils.GetNow())
	return nil
}

func (user *User) Delete() *errors.RestError {
	bookstore_users.Client.Model(&user).Update("DeletedAt", date_utils.GetNow())
	return nil
}
