package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//querystring主要用于GET请求的部分，传输的数据量比较小

func main() {
	r := gin.Default()

	//GET请求127.0.0.1:8080/web?query=zcy
	//GET 请求URL ?后面是querystring参数
	//key-value格式，多个querystring使用&连接
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器发送请求携带的参数querystring
		//方法1：
		//	name := c.Query("query")
		//方法2：如果查不到就使用默认的value
		//	name := c.DefaultQuery("query","hello")

		//方法3：判断是否接受到了
		name, ok := c.GetQuery("query")
		if !ok {
			//取不到
			name = "somebody"
		}
		c.JSON(http.StatusOK,name)
	})

	//GET请求127.0.0.1:8080/web?query=zcy&age=34
	r.GET("/string", func(c *gin.Context) {
		//获取浏览器发送请求携带的参数querystring

		name := c.Query("query")
		age := c.Query("age")

		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})

	r.Run(":8080")
}











