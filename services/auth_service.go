package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
	Logout(ctx context.Context, token string) error
}

type authService struct {
	connections map[string]bool
}

func (a *authService) Login(ctx context.Context, username string, password string) (string, error) {
	if username == "sam" && password == "password" {
		token := uuid.New().String()
		a.connections[token] = true
		return token, nil
	}

	return "", errors.New("invalid username or password")
}

func (a *authService) ValidateToken(ctx context.Context, token string) (bool, error) {
	if _, ok := a.connections[token]; ok {
		return true, nil
	}

	return false, errors.New("invalid token")
}

func (a *authService) Logout(ctx context.Context, token string) error {
	if _, ok := a.connections[token]; ok {
		delete(a.connections, token)
		return nil
	}

	return errors.New("invalid token")
}

// constructor
func NewAuthService() *authService {
	return &authService{
		connections: make(map[string]bool),
	}
}
