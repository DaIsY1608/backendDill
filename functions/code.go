package functions

// func Delete(c *gin.Context)  {
	 
// 		_, CookieError := c.Request.Cookie("dilafGangster")
// 		var shbDelete structs.DeleteProduct
// 		c.ShouldBind(&shbDelete)
// 		var isExist bool = false
// 		if CookieError != nil {
// 			fmt.Println("CookieError: %v\n", CookieError)
// 		} else {
// 			for i, v := range slices.Courses {
// 				if shbDelete.Tittle == "" {
// 					c.JSON(404, "Empty!")
// 				} else if shbDelete.Tittle == v.Tittle {
// 					slices.Courses = append(slices.Courses[:i], slices.Courses[i+1:]...)
// 					c.JSON(200, "Course deleted!")
// 					fmt.Println(slices.Courses)
// 				} else {
// 					isExist = true
// 				}
// 			}
// 			if isExist {
// 				c.JSON(404, "No such course exists!")
// 			}
// 		}
// 	}
gin