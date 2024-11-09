package main

import (
	"github.com/labstack/echo/v4"
	"go-di-sample/application/service"
	"go-di-sample/infrastructure/api/jsonplaceholder"
	"go-di-sample/interface/handlers"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	c := jsonplaceholder.NewClient()
	s := service.NewUserService(c)
	h := handlers.NewUserHandler(s)

	e := echo.New()
	e.GET("/users", h.GetUserList)
	e.GET("/users/:id", h.GetUserById)

	e.Logger.Fatal(e.Start(":1234"))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
