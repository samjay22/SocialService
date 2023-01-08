package services

import (
	"context"
	"github.com/samjay22/SocialService/database"
	"github.com/samjay22/SocialService/structs"
)

type UserService interface {
	GetUserByID(ctx context.Context, userId string) (*structs.UserProfile, error)
	GetUserByEmail(email string) (*structs.UserProfile, error)
	GetUserByUsername(username string) (*structs.UserProfile, error)

	UpdateUser(user *structs.UserProfile) error
	CreateUser(ctx context.Context, userId string) error
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

func (u userService) CreateUser(ctx context.Context, userId string) error {

	newUserProfile := structs.NewUserProfile("1/7/2023", "NONE", "sam@gmail.com", "757-752-0752", "1 house st", "None", "None", 1, "Male")
	u.database.AddUserFromId(ctx, userId, newUserProfile)

	return nil
}

func (u userService) DeleteUser(user *structs.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetUserByID(ctx context.Context, userId string) (*structs.UserProfile, error) {
	return u.database.GetUserByID(ctx, userId)
}

func NewUserService(databaseObject database.Database) *userService {
	return &userService{
		database: databaseObject,
	}
}
