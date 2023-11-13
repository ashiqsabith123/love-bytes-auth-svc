//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/ashiqsabith123/auth-svc/pkg/db"
	"github.com/ashiqsabith123/auth-svc/pkg/repository"
	"github.com/ashiqsabith123/auth-svc/pkg/service"
	"github.com/ashiqsabith123/auth-svc/pkg/usecase"
	"github.com/google/wire"
)

func IntializeService(config config.Config) service.UserService {

	wire.Build(
		db.ConnectToDatabase,
		repository.NewUserRepo,
		usecase.NewUserUsecase,
		service.NewUserService,
	)

	return service.UserService{}

}
