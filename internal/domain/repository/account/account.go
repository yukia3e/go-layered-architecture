package account

import (
	"context"

	"github.com/yukia3e/go-layered-architecture/internal/domain/model/account"
)

type Repository interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) error
	GetAccount(ctx context.Context, userID string) (*account.User, error)
}

type CreateAccountRequest struct {
	UserID     string
	HandleName string
}
