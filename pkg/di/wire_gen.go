// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/ashiqsabith123/auth-svc/pkg/db"
	"github.com/ashiqsabith123/auth-svc/pkg/repository"
	"github.com/ashiqsabith123/auth-svc/pkg/service"
	"github.com/ashiqsabith123/auth-svc/pkg/usecase"
)

// Injectors from wire.go:

func IntializeService(config2 config.Config) service.UserService {
	gormDB := db.ConnectToDatabase(config2)
	userRepo := repository.NewUserRepo(gormDB)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userService := service.NewUserService(userUsecase)
	return userService
}
