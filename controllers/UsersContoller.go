package controllers

import (
	"dilya/collab/data"
	"dilya/collab/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var NewUser models.UserModel
	c.ShouldBindJSON(&NewUser)

	if NewUser.Name == "" || NewUser.Email == "" || NewUser.Age == 0 || NewUser.Password == "" {
		c.JSON(400, gin.H{
			"message": "Empty fields are not allowed",
		})
	} else {
		NewUser.Id = fmt.Sprintf("%v", rand.Intn(10000)*time.Now().Second())
		data.UsersSlice = append(data.UsersSlice, NewUser)
		c.JSON(200, gin.H{
			"message": "Created successfully",
			"data":    NewUser,
		})
	}
}
