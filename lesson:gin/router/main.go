package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":"get",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":"post",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":"put",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":"delete",
		})
	})

	//可以处理所有的方法
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet  :c.JSON(http.StatusOK,gin.H{"method":"GET"})
		case http.MethodPost :c.JSON(http.StatusOK,gin.H{"method":"POST"})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":"any",
		})
	})

	//处理没有定义的路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":"not found",
		})
	})

	//路由组：
	//把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status":"get"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status":"xx"})
		})
		videoGroup.GET("/haha", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status":"haha"})
		})
	}


	r.Run(":8080")
}














