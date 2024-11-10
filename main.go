package main

import (
	"github.com/labstack/echo/v4"
	"go-di-sample/di"
)

func main() {
	h := di.InitializeUserHandler()

	e := echo.New()
	e.GET("/users", h.GetUserList)
	e.GET("/users/:id", h.GetUserById)

	e.Logger.Fatal(e.Start(":1234"))
}
