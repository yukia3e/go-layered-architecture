package grpc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"

	accountUsecase "github.com/yukia3e/go-layered-architecture/internal/feature/account/usecase"
	accountv1 "github.com/yukia3e/go-layered-architecture/internal/ui/grpc/gen/account/v1"
)

type Server interface {
	CreateAccount(context.Context, *connect.Request[accountv1.CreateAccountRequest]) (*connect.Response[accountv1.CreateAccountResponse], error)
	GetAccount(ctx context.Context, req *connect.Request[accountv1.GetAccountRequest]) (*connect.Response[accountv1.GetAccountResponse], error)
}

type server struct {
	usecase accountUsecase.Usecase
}

func NewServer(usecase accountUsecase.Usecase) Server {
	return &server{
		usecase: usecase,
	}
}

func (s *server) CreateAccount(ctx context.Context, req *connect.Request[accountv1.CreateAccountRequest]) (*connect.Response[accountv1.CreateAccountResponse], error) {
	log.Debug().Msgf("Request headers: %s", req.Header())

	userID, err := s.usecase.CreateAccount(ctx, req.Msg.HandleName)
	if err != nil {
		log.Error().AnErr("error", err).Msg("account.Server.CreateAccount: failed to a create account")
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&accountv1.CreateAccountResponse{
		UserId: userID,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (s *server) GetAccount(ctx context.Context, req *connect.Request[accountv1.GetAccountRequest]) (*connect.Response[accountv1.GetAccountResponse], error) {
	user, err := s.usecase.GetAccount(ctx, req.Msg.UserId)
	if err != nil {
		log.Error().AnErr("error", err).Msg("account.Server.GetAccount: failed to a get account")
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&accountv1.GetAccountResponse{
		UserId:     req.Msg.UserId,
		HandleName: user.HandleName,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}
