package database

import (
	"errors"
	"fmt"
	"github.com/samjay22/SocialService/structs"
)

type Database interface {
	GetUserByID(userID string) (*structs.UserProfile, error)
	AddUserFromId(userID string, userProfile *structs.UserProfile) error
}

type mockDatabase struct {
	Users map[string]*structs.UserProfile
}

func NewMockDatabase() *mockDatabase {
	userBase := map[string]*structs.UserProfile{
		"1": {
			Username: "1",
		},
		"2": {
			Username: "2",
		},
		"3": {
			Username: "3",
		},
	}

	return &mockDatabase{
		Users: userBase,
	}
}

func (d mockDatabase) GetUserByID(userID string) (*structs.UserProfile, error) {
	user, ok := d.Users[userID]

	fmt.Println(d.Users[userID], userID)

	if !ok {
		return nil, errors.New("Error not found!")
	}

	return user, nil
}

func (d mockDatabase) AddUserFromId(userID string, userProfile *structs.UserProfile) error {
	d.Users[userID] = userProfile
	return nil
}
