package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/samjay22/SocialService/structs"
)

type Database interface {
	GetUserByID(ctx context.Context, userID string) (*structs.UserProfile, error)
	GetUserByEmail(email string) (*structs.UserProfile, error)
	GetUserByUsername(username string) (*structs.UserProfile, error)
	UpdateUser(user *structs.UserProfile) error
	DeleteUser(user *structs.UserProfile) error
	AddUserFromID(ctx context.Context, userID string, userProfile *structs.UserProfile) error
}

type mockDatabase struct {
	redisStore *redis.Client
}

func NewMockDatabase(redisURL string) (*mockDatabase, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	redisStore := redis.NewClient(options)

	return &mockDatabase{
		redisStore: redisStore,
	}, nil
}

func (d *mockDatabase) GetUserByID(ctx context.Context, userID string) (*structs.UserProfile, error) {
	userProfileKey := userProfileKey(userID)
	userProfileData, err := d.redisStore.Get(ctx, userProfileKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	userProfile := &structs.UserProfile{}
	if err := json.Unmarshal(userProfileData, userProfile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user profile: %w", err)
	}

	return userProfile, nil
}

func (d *mockDatabase) AddUserFromID(ctx context.Context, userID string, userProfile *structs.UserProfile) error {
	userProfileKey := userProfileKey(userID)
	userProfileData, err := json.Marshal(userProfile)
	if err != nil {
		return fmt.Errorf("failed to marshal user profile: %w", err)
	}

	if err := d.redisStore.Set(ctx, userProfileKey, userProfileData, 0).Err(); err != nil {
		return fmt.Errorf("failed to set user profile: %w", err)
	}

	return nil
}

func (d *mockDatabase) GetUserByEmail(email string) (*structs.UserProfile, error) {
	userProfileKey := userProfileKey(email)
	userProfileData, err := d.redisStore.Get(context.Background(), userProfileKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	userProfile := &structs.UserProfile{}
	if err := json.Unmarshal(userProfileData, userProfile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user profile: %w", err)
	}

	return userProfile, nil
}

func (d *mockDatabase) GetUserByUsername(username string) (*structs.UserProfile, error) {
	userProfileKey := userProfileKey(username)
	userProfileData, err := d.redisStore.Get(context.Background(), userProfileKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	userProfile := &structs.UserProfile{}
	if err := json.Unmarshal(userProfileData, userProfile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user profile: %w", err)
	}

	return userProfile, nil
}

func (d *mockDatabase) UpdateUser(user *structs.UserProfile) error {
	userProfileKey := userProfileKey(user.ID)
	userProfileData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user profile: %w", err)
	}

	if err := d.redisStore.Set(context.Background(), userProfileKey, userProfileData, 0).Err(); err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}

	return nil
}

func (d *mockDatabase) DeleteUser(user *structs.UserProfile) error {
	userProfileKey := userProfileKey(user.ID)

	if err := d.redisStore.Del(context.Background(), userProfileKey).Err(); err != nil {
		return fmt.Errorf("failed to delete user profile: %w", err)
	}

	return nil
}

func userProfileKey(userID string) string {
	return fmt.Sprintf("%s_UserProfile", userID)
}
