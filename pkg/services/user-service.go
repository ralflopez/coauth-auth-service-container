package services

import (
	"coauth/pkg/models"
	"coauth/pkg/utils"

	"github.com/google/uuid"
)

var users []*models.User = []*models.User{}

func GetUsers() []*models.User {
	return users
}

func GetUser(id string) *models.User {
	return &models.User{}
}

func CreateUser(u *models.User) error {
	var err error

	u.Id = uuid.NewString()
	if u.Role == "" {
		u.Role = models.Member
	}
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	users = append(users, u)
	
	return nil
}

func UpdateUser(u *models.User) error {
	return nil
}

func DelteUser(u *models.User) error {
	return nil
}
