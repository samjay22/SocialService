package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/samjay22/SocialService/database"
	"github.com/samjay22/SocialService/structs"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID int64) (*structs.UserProfile, error)
	CreateUser(ctx context.Context, userID *structs.CreateUserProfileRequest) error
	AddUserAsFriend(ctx context.Context, userID int64, friendID int64) error
}

type userService struct {
	database database.Database
}

func (u *userService) CreateUser(ctx context.Context, newUserProfile *structs.CreateUserProfileRequest) error {
	fmt.Println(newUserProfile.Name, newUserProfile.Age, newUserProfile.JoinDate)
	if newUserProfile.Name == "" || newUserProfile.Age == 0 || newUserProfile.JoinDate == "" {
		return errors.New("invalid user profile")
	}

	profileToAdd := structs.NewUserProfile(newUserProfile.ID, newUserProfile.JoinDate, newUserProfile.Name, newUserProfile.Age)
	return u.database.AddUserFromID(ctx, newUserProfile.ID, profileToAdd)
}

func (u *userService) GetUserByID(ctx context.Context, userID int64) (*structs.UserProfile, error) {
	return u.database.GetUserByID(ctx, userID)
}

func (u *userService) AddUserAsFriend(ctx context.Context, userID int64, friendID int64) error {
	if userID == friendID {
		return errors.New("cannot add self as friend")
	}

	user, err := u.database.GetUserByID(ctx, userID)
	if user == nil {
		return errors.New("user does not exist")
	} else if err != nil {
		return err
	}

	friend, err := u.database.GetUserByID(ctx, friendID)
	if friend == nil {
		return errors.New("friend does not exist")
	} else if err != nil {
		return err
	}

	for _, friend := range user.Friends {
		if friend == friendID {
			return errors.New("friend already exists")
		}
	}
	user.Friends = append(user.Friends, friend.ID)

	err = u.database.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(databaseObject database.Database) *userService {
	return &userService{
		database: databaseObject,
	}
}
