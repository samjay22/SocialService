package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/samjay22/SocialService/structs"
)

type Database interface {
	GetUserByID(ctx context.Context, userID int64) (*structs.UserProfile, error)
	AddUserFromID(ctx context.Context, userID int64, userProfile *structs.UserProfile) error
	UpdateUser(context.Context, *structs.UserProfile) error
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

func (d *mockDatabase) GetUserByID(ctx context.Context, userID int64) (*structs.UserProfile, error) {
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

// Update user profile
func (d *mockDatabase) UpdateUser(ctx context.Context, userProfile *structs.UserProfile) error {
	userProfileKey := userProfileKey(userProfile.ID)
	userProfileData, err := json.Marshal(userProfile)
	if err != nil {
		return fmt.Errorf("failed to marshal user profile: %w", err)
	}

	if err := d.redisStore.Set(ctx, userProfileKey, userProfileData, 0).Err(); err != nil {
		return fmt.Errorf("failed to set user profile: %w", err)
	}

	return nil
}

func (d *mockDatabase) AddUserFromID(ctx context.Context, userID int64, userProfile *structs.UserProfile) error {
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

func userProfileKey(userID int64) string {
	return fmt.Sprint(userID)
}
