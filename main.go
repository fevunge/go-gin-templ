package main

import (
	"github.com/fevunge/let-go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controller.CreateUser(router)

	router.Run(":8080")
}
