package main

import (
	"dilya/collab/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/users/add", controllers.AddUser)

	router.Run(":4004")
}
