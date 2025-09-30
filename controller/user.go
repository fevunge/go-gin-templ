// Package controller:: users
package controller

import "github.com/labstack/echo/v4"

func CreateUser(e *echo.Echo) {
	user := e.Group("/users")
	user.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		response := "Create User: " + name + " with email: " + email
		return c.String(201, response)
	})
}
