package account

import (
	"context"

	accountModel "github.com/yukia3e/go-layered-architecture/internal/domain/model/account"
	"github.com/yukia3e/go-layered-architecture/internal/domain/repository/account"
)

type memory struct {
	users map[string]accountModel.User
}

func NewMemory() account.Repository {
	return &memory{
		users: map[string]accountModel.User{},
	}
}

func (m *memory) CreateAccount(ctx context.Context, req account.CreateAccountRequest) error {
	m.users[req.UserID] = accountModel.User{
		HandleName: req.HandleName,
	}
	return nil
}

func (m *memory) GetAccount(ctx context.Context, userID string) (*accountModel.User, error) {
	user := m.users[userID]
	return &user, nil
}
