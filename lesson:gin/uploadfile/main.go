package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/index",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//上传一个文件的例子：
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB

	r.POST("/upload",func(c *gin.Context) {
		//从请求中读取文件
		//此处的文件名字由这里来：<input type="file" name="f1">
		file,err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
		} else {
			//把读取到的文件存放到本地
			dst := fmt.Sprintf("./%s",file.Filename)
			//方法2：dst := path.Join("./", file.Filename)
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"status":"ok",
			})
		}

	})

	//上传多个文件：
	r.POST("/multifiles", func(c *gin.Context) {
		form,_ := c.MultipartForm()
		files := form.File["file"]

		for _,file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":"ok",
		})
	})

	r.Run(":8080")
}












