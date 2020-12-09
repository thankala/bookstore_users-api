package users

import (
	"encoding/json"
	"time"
)

type PublicUser struct {
	Id        int64   `json:"Id"`
	Status    string `json:"status"`
	CreatedAt time.Time
}

type PrivateUser struct {
	Id        int64   `json:"Id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:        user.Id,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}
