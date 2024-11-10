//go:build wireinject

package di

import (
	"github.com/google/wire"
	"go-di-sample/application/service"
	"go-di-sample/infrastructure/api/jsonplaceholder"
	"go-di-sample/interface/handler"
)

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(
		jsonplaceholder.NewClient,
		service.NewUserService,
		handler.NewUserHandler,
	)

	return &handler.UserHandler{}
}
