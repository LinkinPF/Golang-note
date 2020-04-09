package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//跳转到百度
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
	})

	//这个是跳转到别的路由处理函数
	r.GET("/a", func(c *gin.Context) {
		//修改请求的URL
		c.Request.URL.Path = "/test"
		//延续后续的处理
		r.HandleContext(c)
	})

	r.Run(":8080")
}
