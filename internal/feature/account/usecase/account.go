package usecase

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/pkg/errors"
	"github.com/yukia3e/go-layered-architecture/internal/feature/account/domain/model"
	accountRepository "github.com/yukia3e/go-layered-architecture/internal/feature/account/domain/repository"
)

type Usecase interface {
	CreateAccount(ctx context.Context, handleName string) (string, error)
	GetAccount(ctx context.Context, userID string) (*model.User, error)
}

type usecase struct {
	repository accountRepository.Repository
}

func NewUsecase(repository accountRepository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateAccount(ctx context.Context, handleName string) (string, error) {
	log.Debug().Msg("account.Usecase.CreateAccount: called")

	// TODO: ssoからのUserID取得
	userID := "UserA"

	if err := u.repository.CreateAccount(ctx,
		accountRepository.CreateAccountRequest{
			UserID:     userID,
			HandleName: handleName,
		},
	); err != nil {
		return "", errors.Wrap(err, "account.Usecase.CreateAccount")
	}

	return userID, nil
}

func (u *usecase) GetAccount(ctx context.Context, userID string) (*model.User, error) {
	log.Debug().Msg("account.Usecase.GetAccount: called")

	user, err := u.repository.GetAccount(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "account.Usecase.GetAccount")
	}

	return user, nil
}
