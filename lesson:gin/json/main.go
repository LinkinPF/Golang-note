package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {

		//方法1：使用map
		//data := map[string]interface{}{
		//	"name" : "zcy",
		//	"age" : 3,
		//}

		//方法2：使用gin.H:   type H map[string]interface{}   gin框架帮我们定义了
		data := gin.H{
			"name" : "zcy",
			"age" : 3,
		}

		c.JSON(http.StatusOK,data)
	})

	r.GET("/jsonstruct", func(c *gin.Context) {
		//方法3：定义结构体
		type msg struct {
			Name string	`json:"name"`	//如果写成小写字母开头，那么就不能传输了,如果前端需要，就使用tag
			Age int	`json:"age"`
		}
		var data msg
		data.Name = "haha"
		data.Age = 34

		data2 := msg{
			"ko",
			23,
		}

		data3 := gin.H{
			"data" : data,
			"data2" : data2,
		}

		c.JSON(http.StatusOK,data3)
	})

	r.Run(":8080")
}













