package services

import (
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/utils/crypto_utils"
	"github.com/thankala/bookstore_users-api/utils/errors"
)

var (
	UsersService UserService = &usersService{}
)

type UserService interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(uint) (*users.User, *errors.RestError)
	UpdateUser(users.User) (*users.User, *errors.RestError)
	DeleteUser(uint) *errors.RestError
	SearchUser(string) (users.Users, *errors.RestError)
	LoginUser(users.User) (*users.User, *errors.RestError)
}

type usersService struct {
}

func (usersService *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if validateError := user.Validate(); validateError != nil {
		return nil, validateError
	}
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (usersService *usersService) GetUser(userId uint) (*users.User, *errors.RestError) {
	user := users.User{ID: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (usersService *usersService) UpdateUser(user users.User) (*users.User, *errors.RestError) {
	current, err := usersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" {
		current.LastName = user.LastName
	}
	if user.Email != "" {
		current.Email = user.Email
	}
	if validateError := current.Validate(); validateError != nil {
		return nil, validateError
	}
	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (usersService *usersService) DeleteUser(userId uint) *errors.RestError {
	user := users.User{ID: userId}
	return user.Delete()
}

func (usersService *usersService) SearchUser(status string) (users.Users, *errors.RestError) {
	users := users.Users{}
	if err := users.FindByStatus(status); err != nil {
		return nil, err
	}
	return users, nil
}

func (usersService *usersService) LoginUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return &user, nil
}
