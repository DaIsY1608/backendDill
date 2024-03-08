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

func DeleteUser(c *gin.Context) {
	var shbDelete models.DeleteUser
	c.ShouldBind(&shbDelete)

	var isExist bool = false

	for i, v := range data.UsersSlice {
		if shbDelete.Name == "" {
			c.JSON(404, "Empty!")
		} else if shbDelete.Name == v.Name {
			data.UsersSlice = append(data.UsersSlice[:i], data.UsersSlice[i+1:]...)
			c.JSON(200, "Course deleted!")
			fmt.Println(data.UsersSlice)
		} else {
			isExist = true
		}
	}
	if isExist {
		c.JSON(404, "No such User exists!")
	}

}