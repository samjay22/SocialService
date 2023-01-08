package services

import (
	"github.com/samjay22/SocialService/database"
	"github.com/samjay22/SocialService/structs"
)

type UserService interface {
	GetUserByID(username string) (*structs.UserProfile, error)
	GetUserByEmail(email string) (*structs.UserProfile, error)
	GetUserByUsername(username string) (*structs.UserProfile, error)

	UpdateUser(user *structs.UserProfile) error
	CreateUser(user *structs.UserProfile) error
	DeleteUser(user *structs.UserProfile) error
}

type userService struct {
	database database.Database
}

func (u userService) GetUserByEmail(email string) (*structs.UserProfile, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetUserByUsername(username string) (*structs.UserProfile, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) UpdateUser(user *structs.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) CreateUser(user *structs.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) DeleteUser(user *structs.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetUserByID(userId string) (*structs.UserProfile, error) {
	return u.database.GetUserByID(userId)
}

func NewUserService(databaseObject database.Database) *userService {
	return &userService{
		database: databaseObject,
	}
}
