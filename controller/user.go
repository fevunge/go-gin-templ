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
		username := c.FormValue("username")
		password := c.FormValue("password")

		response, err := model.CreateUser(name, username, password)
		if err != nil {
			errorMessage := map[string]string{"error": err.Error()}
			return c.JSON(500, errorMessage)
		}
		return c.JSON(201, response)
	})
}

func GetProfile(e *echo.Echo) {
	user := e.Group("/users")
	user.GET("/profile", func(c echo.Context) error {
		username := c.QueryParam("username")

		response, err := model.GetProfile(username)
		if err != nil {
			errorMessage := map[string]string{"error": err.Error()}
			return c.JSON(500, errorMessage)
		}
		return c.JSON(200, response)
	})
}