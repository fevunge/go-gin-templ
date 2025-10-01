// Package controller:: users
package controller

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(r *gin.Engine) {
	r.GET("/create", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":     "John Doe",
			"email":    "Johndoe@email.com",
			"age":      30,
			"is_admin": false,
			"me":       c.ClientIP(),
		})
	})
}

func GetName(r *gin.Engine) {
	data := Param{}

	r.GET("/get", func(c *gin.Context) {
		c.BindJSON(&data)
		c.JSON(200, gin.H{
			"query": data.Fields,
		})
	})
}
