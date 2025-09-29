// Package controller:: users
package controller

import "github.com/labstack/echo/v4"

func CreateUser(e *echo.Echo) {
	e.POST("/users", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		response := "Create User: " + name + " with email: " + email
		return c.String(201, response)
	})
}
