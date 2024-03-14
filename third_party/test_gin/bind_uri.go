package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}

// 测试：
/*
➜  Xray-windows-64-B curl http://localhost:8088/sunbin/123456
{"msg":"Key: 'Person.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"}%                                        ➜  Xray-windows-64-B curl http://localhost:8088/sunbin/987fbc97-4bed-5078-9f07-9141ba07c9f3
{"name":"sunbin","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3"}%
*/
