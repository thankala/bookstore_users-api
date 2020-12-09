package users

import (
	"github.com/thankala/bookstore_users-api/utils/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	StatusActive = "active"
)

type User struct {
	gorm.Model
	Id        int64    `gorm:"primaryKey" json:"Id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Status    string `json:"status"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Users []User

func (user User) Validate() *errors.RestError {
	if user.Password == "" {
		return errors.NewBadRequestError("Password field empty")
	}
	if user.Email == "" {
		return errors.NewBadRequestError("Email field empty")
	}
	if user.FirstName == "" {
		return errors.NewBadRequestError("First name field empty")
	}
	if user.LastName == "" {
		return errors.NewBadRequestError("Last name field empty")
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	return nil
}
