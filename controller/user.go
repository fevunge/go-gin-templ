// Package controller:: users
package controller

import (
	"net/http"
	"strconv"

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
	data := []string{"Banana", "Apple", "Orange"}

	r.GET("/get/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "ID must be a number",
				"result":  nil,
			})
			return
		}
		if id > len(data) || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Index out of range 1-" + strconv.Itoa(len(data)),
				"result":  nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Everything is OK",
			"result":  data[id-1],
		})
	})
}
