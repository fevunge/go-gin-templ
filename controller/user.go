// Package controller:: users
package controller

import (
	"github.com/fevunge/let-go/model"
	"github.com/labstack/echo/v4"
)

func CreateUser(e *echo.Echo) {
	user := e.Group("/users")
	user.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		username := c.FormValue("username")
		password := c.FormValue("password")

		model.CreateUser(name, username, password)

		response := "Create User: " + name + " with email: " + email + username + " and password: " + password
		return c.String(201, response)
	})
}
