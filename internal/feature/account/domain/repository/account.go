package repository

import (
	"context"

	"github.com/yukia3e/go-layered-architecture/internal/feature/account/domain/model"
)

type Repository interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) error
	GetAccount(ctx context.Context, userID string) (*model.User, error)
}

type CreateAccountRequest struct {
	UserID     string
	HandleName string
}
