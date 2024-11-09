package main

import (
	"github.com/labstack/echo/v4"
	"go-di-sample/application/service"
	"go-di-sample/infrastructure/api/jsonplaceholder"
	"go-di-sample/interface/handler"
)

func main() {
	c := jsonplaceholder.NewClient()
	s := service.NewUserService(c)
	h := handler.NewUserHandler(s)

	e := echo.New()
	e.GET("/users", h.GetUserList)
	e.GET("/users/:id", h.GetUserById)

	e.Logger.Fatal(e.Start(":1234"))
}
