package services

import (
	"context"

	"github.com/samjay22/SocialService/database"
	"github.com/samjay22/SocialService/structs"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID string) (*structs.UserProfile, error)
	GetUserByEmail(email string) (*structs.UserProfile, error)
	GetUserByUsername(username string) (*structs.UserProfile, error)

	UpdateUser(user *structs.UserProfile) error
	CreateUser(ctx context.Context, userID string) error
	DeleteUser(user *structs.UserProfile) error
}

type userService struct {
	database database.Database
}

func (u *userService) GetUserByEmail(email string) (*structs.UserProfile, error) {
	return u.database.GetUserByEmail(email)
}

func (u *userService) GetUserByUsername(username string) (*structs.UserProfile, error) {
	return u.database.GetUserByUsername(username)
}

func (u *userService) UpdateUser(user *structs.UserProfile) error {
	return u.database.UpdateUser(user)
}

func (u *userService) CreateUser(ctx context.Context, userID string) error {
	newUserProfile := structs.NewUserProfile(userID, "1/7/2023", "Tom", "000-000-0000", 16)
	return u.database.AddUserFromID(ctx, userID, newUserProfile)
}

func (u *userService) DeleteUser(user *structs.UserProfile) error {
	return u.database.DeleteUser(user)
}

func (u *userService) GetUserByID(ctx context.Context, userID string) (*structs.UserProfile, error) {
	return u.database.GetUserByID(ctx, userID)
}

func NewUserService(databaseObject database.Database) *userService {
	return &userService{
		database: databaseObject,
	}
}
