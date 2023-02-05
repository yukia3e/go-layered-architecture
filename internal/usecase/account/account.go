package account

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/pkg/errors"
	accountModel "github.com/yukia3e/go-layered-architecture/internal/domain/model/account"
	"github.com/yukia3e/go-layered-architecture/internal/domain/repository/account"
)

type Usecase interface {
	CreateAccount(ctx context.Context, handleName string) (string, error)
	GetAccount(ctx context.Context, userID string) (*accountModel.User, error)
}

type usecase struct {
	repository account.Repository
}

func NewUsecase(repository account.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateAccount(ctx context.Context, handleName string) (string, error) {
	log.Debug().Msg("account.Usecase.CreateAccount: called")

	// TODO: ssoからのUserID取得
	userID := "UserA"

	if err := u.repository.CreateAccount(ctx,
		account.CreateAccountRequest{
			UserID:     userID,
			HandleName: handleName,
		},
	); err != nil {
		return "", errors.Wrap(err, "account.Usecase.CreateAccount")
	}

	return userID, nil
}

func (u *usecase) GetAccount(ctx context.Context, userID string) (*accountModel.User, error) {
	log.Debug().Msg("account.Usecase.GetAccount: called")

	user, err := u.repository.GetAccount(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "account.Usecase.GetAccount")
	}

	return user, nil
}
