package users

import (
	"github.com/thankala/bookstore_users-api/utils/errors"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey; autoIncrement" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique; not null" json:"email"`
}

func (user User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" || user.LastName == "" || user.FirstName == "" {
		return errors.NewBadRequestError("Invalid input")
	}
	return nil
}
