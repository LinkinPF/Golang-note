package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


//form主要用于POST请求，发送的数据量比较大了
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	//这个请求只是传回给浏览器登录界面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//
	r.POST("/login", func(c *gin.Context) {
		//获取表单传输的数据
		//方法1：PostForm
		//	这的key值要和html中的对应：<input type="text" name="username" id="username">中的name
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//c.HTML(http.StatusOK, "index.html", gin.H{
		//	"Name" : username,
		//	"Word" : password,
		//})
		//方法2：
		//	使用GetPostForm

		//浏览器什么都不输入，ok也不是false，此时接收到的就是空字符串。也是true
		username,ok := c.GetPostForm("username")
		if !ok {
			username = "none"
		}
		password,ok := c.GetPostForm("password")
		if !ok {
			password = "none"

		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name" : username,
			"Word" : password,
		})
	})


	r.Run(":8080")


}
























