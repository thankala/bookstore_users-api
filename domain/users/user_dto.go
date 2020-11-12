package users

import (
	"github.com/thankala/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" || user.DateCreated == "" || user.LastName == "" || user.FirstName == "" || user.Id == 0 {
		return errors.NewBadRequestError("Invalid input")
	}
	return nil
}
