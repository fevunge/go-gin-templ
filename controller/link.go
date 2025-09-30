// Package controller handles the HTTP requests and responses for link-related operations.
package controller

import "github.com/labstack/echo/v4"

func AddLink(e *echo.Echo) error {
	link := e.Group("/links")
	link.POST("/", func(c echo.Context) error {
		c.JSON(201, ["OK"])
	})
	return nil
}
