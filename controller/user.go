// Package controller:: users
package controller

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":     "John Doe",
			"email":    "Johndoe@email.com",
			"age":      30,
			"is_admin": false,
		})
	})
}
