package main

import (
	"github.com/fevunge/let-go/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to LetGo API")
	})

	controller.CreateUser(e)

	e.Logger.Fatal(e.Start(":1323"))
}
