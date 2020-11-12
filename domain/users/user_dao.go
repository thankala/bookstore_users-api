package users

import (
	"fmt"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"time"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found",user.Id))
	}
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists",user.Id))
	}
	now := time.Now()
	user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	usersDB[user.Id] = user
	return nil
}
