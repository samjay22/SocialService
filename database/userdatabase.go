package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"github.com/samjay22/SocialService/structs"
	"os"
)

type Database interface {
	GetUserByID(ctx context.Context, userId string) (*structs.UserProfile, error)
	AddUserFromId(ctx context.Context, userID string, userProfile *structs.UserProfile) error
}

type mockDatabase struct {
	redisStore *redis.Client
}

func NewMockDatabase() *mockDatabase {

	redisStore := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	return &mockDatabase{
		redisStore: redisStore,
	}
}

func (d mockDatabase) GetUserByID(ctx context.Context, userId string) (*structs.UserProfile, error) {
	userProfileKey := userProfileKey(userId)
	userProfile := d.redisStore.Get(ctx, userProfileKey)

	if userProfile == nil {
		return nil, errors.New("User not found")
	}

	userProfileData, err := userProfile.Bytes()
	if err != nil {
		return nil, err
	}

	userProfileStruct := &structs.UserProfile{}
	err = jsoniter.Unmarshal(userProfileData, userProfileStruct)
	if err != nil {
		return nil, err
	}

	return userProfileStruct, nil
}

func (d mockDatabase) AddUserFromId(ctx context.Context, userID string, userProfile *structs.UserProfile) error {

	userProfileKey := userProfileKey(userID)
	userProfileAsBytes, err := convertUserProfileToByteSlice(userProfile)
	if err != nil {
		return err
	}

	requestStatus := d.redisStore.Set(context.Background(), userProfileKey, userProfileAsBytes, 0)
	if requestStatus.Err() != nil {
		return requestStatus.Err()
	}

	return nil
}

func userProfileKey(userId string) string {
	return fmt.Sprintf("%s_UserProfile")
}

func convertUserProfileToByteSlice(userProfile *structs.UserProfile) ([]byte, error) {
	return jsoniter.Marshal(userProfile)
}
