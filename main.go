package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register")
	r.POST("/list")
	r.POST("/delete", )
	r.Run(":2020")

}
