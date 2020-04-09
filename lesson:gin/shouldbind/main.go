package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//相让别人能够拿到结构体中的字段，需要大写开头
//另外要想使用ShouldBind函数，需要告诉它会获取什么数据，使用form tag
type UserInfo struct {
	Username string	`form:"username" json:"user"`
	Password string	`form:"password" json:"pwd"`
}

//参数绑定：前端数据和后端结构体做一个绑定
func main() {
	r := gin.Default()

	//这是是querystring的方式
	//	访问：http://127.0.0.1:8080/user?username=zcy&password=123
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		//想让前端传来的数据和后端结构体绑定,注意传递指针
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		} else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK, gin.H{
				"ststus" : "ok",
			})
		}
	})

	//这是form表单的方式
	//使用postman在body中填写
	r.POST("/post", func(c *gin.Context) {
		var u UserInfo
		//想让前端传来的数据和后端结构体绑定,注意传递指针
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		} else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK, gin.H{
				"ststus" : "ok",
			})
		}
	})

	//这里处理前端发过来的json的数据
	r.POST("/json", func(c *gin.Context) {
		var u UserInfo
		//想让前端传来的数据和后端结构体绑定,注意传递指针
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		} else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK, gin.H{
				"ststus" : "ok",
			})
		}
	})

	//注意前面不管是哪种方式，后端的代码都一模一样，说明这个shouldbin牛逼
	r.Run(":8080")
}
