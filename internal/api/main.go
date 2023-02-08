package api

import (
	"crypto/tls"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/yukia3e/go-layered-architecture/internal/config"
	"github.com/yukia3e/go-layered-architecture/internal/ui/grpc/gen/account/v1/accountv1connect"

	accountMemory "github.com/yukia3e/go-layered-architecture/internal/feature/account/infrastructure/memory"
	accountGrpc "github.com/yukia3e/go-layered-architecture/internal/feature/account/ui/grpc"
	accountUsecase "github.com/yukia3e/go-layered-architecture/internal/feature/account/usecase"
	health "github.com/yukia3e/go-layered-architecture/internal/feature/health/ui/http"
)

func Run() error {
	r := gin.Default()

	// health
	r.GET("/health", gin.WrapF(health.HandlerFunc))

	// account
	accountPath, accountHandler := accountv1connect.NewAccountServiceHandler(
		accountGrpc.NewServer(
			accountUsecase.NewUsecase(
				accountMemory.NewMemory(),
			),
		),
	)
	r.Any(accountPath+"*rpc", gin.WrapH(accountHandler))

	cert, err := config.GetServerTLSCredential()
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to get server tls credential")
	}
	endlessServer := endless.NewServer(":8080", r)
	if cert != nil {
		endlessServer.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{*cert},
		}
	} else {
		r.UseH2C = true
	}

	err = endlessServer.ListenAndServe()
	log.Debug().Msg(err.Error())

	return nil
}
