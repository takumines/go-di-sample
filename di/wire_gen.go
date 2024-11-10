// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"go-di-sample/application/service"
	"go-di-sample/infrastructure/api/jsonplaceholder"
	"go-di-sample/interface/handler"
)

// Injectors from wire.go:

func InitializeUserHandler() *handler.UserHandler {
	userRepository := jsonplaceholder.NewClient()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}
