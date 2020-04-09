package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//PATH参数、URI参数
func main() {
	r := gin.Default()

	//注意路径冲突的问题，如果这里改成/:name/:age，那么就会和下面的GET方法的路径冲突
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")		//注意这里的值还是一个字符串的类型
		c.JSON(http.StatusOK, gin.H{
			"name" : name,
			"age" : age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year" : year,
			"month" : month,
		})
	})

	r.Run(":8080")
}
