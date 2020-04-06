package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//这个main函数是gin的快速入门，返回一个状态码，和一个map
//func main() {
//	//创建一个默认的引擎
//	r := gin.Default()
//
//	//get请求方式访问/hello
//	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
//	r.GET("/hello", func(c *gin.Context) {
//		// c.JSON：返回JSON格式的数据
//		c.JSON(http.StatusOK, gin.H{
//			"message": "Hello world!",
//		})
//	})
//
//	//启动HTTP服务，默认地址是127.0.0.1:8080启动服务
//	r.Run(":8080")
//
//}


//这个main函数是gin的HTML的渲染,对一个HTML页面进行渲染
//func main() {
//	//创建一个默认的引擎
//	r := gin.Default()
//
//	//定义模板
//	//模板解析
//	r.LoadHTMLFiles("templates/index.tmpl")
//
//	//get请求方式访问/hello
//	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
//	//这个匿名函数返回一个模板文件，并进行渲染
//	r.GET("/index", func(c *gin.Context) {
//		//HTTP请求
//		//模板渲染
//		c.HTML(http.StatusOK, "index.tmpl", gin.H{
//			//这里的 gin.H 就是一个map：type H map[string]interface{}
//			//这里就在 {{.title}} 模板中使用
//			"title" : "zcy",
//		})
//	})
//
//	//启动HTTP服务，默认地址是127.0.0.1:9090启动服务
//	r.Run(":8080")
//}

//这个main函数是gin的HTML的渲染,对多个HTML页面进行渲染
//func main() {
//	//创建一个默认的引擎
//	r := gin.Default()
//
//	//定义模板
//	//模板解析
//	// 		这个就麻烦了：r.LoadHTMLFiles("templates/index.tmpl")
//	r.LoadHTMLGlob("templates/**/*")
//
//	//两个文件就需要两个路由去使用它
//	r.GET("/posts/index", func(c *gin.Context) {
//		//HTTP请求
//		//模板渲染
//		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
//			//这里的 gin.H 就是一个map：type H map[string]interface{}
//			//这里就在 {{.title}} 模板中使用
//			"title" : "posts/index",
//		})
//	})
//
//	r.GET("/users/index", func(c *gin.Context) {
//		//HTTP请求
//		//模板渲染
//		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
//			//这里的 gin.H 就是一个map：type H map[string]interface{}
//			//这里就在 {{.title}} 模板中使用
//			"title" : "users/index",
//		})
//	})
//
//	//启动HTTP服务，默认地址是127.0.0.1:9090启动服务
//	r.Run(":8080")
//}


//这个main函数是自定义函数的
//func main() {
//	r := gin.Default()
//	//定义自定义函数，用于模板中的{{.title | safe}}
//	r.SetFuncMap(template.FuncMap{
//		"safe" : func(str string) template.HTML {
//			return template.HTML(str)
//		},
//	})
//
//
//	//定义模板
//	//解析模板
//	// 		这个就麻烦了：r.LoadHTMLFiles("templates/index.tmpl")
//	r.LoadHTMLGlob("templates/**/*")
//
//	//两个文件就需要两个路由去使用它
//	r.GET("/posts/index", func(c *gin.Context) {
//		//HTTP请求
//		//模板渲染
//		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
//			//这里的 gin.H 就是一个map：type H map[string]interface{}
//			//这里就在 {{.title}} 模板中使用
//			"title" : "posts/index",
//		})
//	})
//
//	r.GET("/users/index", func(c *gin.Context) {
//		//HTTP请求
//		//模板渲染
//		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
//			//这里的 gin.H 就是一个map：type H map[string]interface{}
//			//这里就在 {{.title}} 模板中使用
//			"title" : "users/index",
//		})
//	})
//
//	//启动HTTP服务，默认地址是127.0.0.1:9090启动服务
//	r.Run(":8080")
//
//}



//静态文件：html上用到的样式文件：.css，js文件，图片
//在templates/users/index.tmpl中引用statics目录下的css文件：
//		注意这里访问网址会有两次请求，第一次返回html文件，
//		第二次会检索到引用了<link rel="stylesheet" href="/xxx/index.css">，
//		就会再发送一次请求，要得到/xxx/index.css，
//		这个时候有了r.Static("/xxx","./statics")，但是要返回/xxx的文件，都需要到./statics去找
func main() {
	r := gin.Default()
	//定义自定义函数，用于模板中的{{.title | safe}}
	r.SetFuncMap(template.FuncMap{
		"safe" : func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//加载静态文件
	//		浏览器要找/xxx有关的都回去找./statics下的文件
	r.Static("/xxx","./statics")

	//定义模板
	//解析模板
	// 		这个就麻烦了：r.LoadHTMLFiles("templates/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")

	//两个文件就需要两个路由去使用它
	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		//模板渲染
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			//这里的 gin.H 就是一个map：type H map[string]interface{}
			//这里就在 {{.title}} 模板中使用
			"title" : "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		//模板渲染
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			//这里的 gin.H 就是一个map：type H map[string]interface{}
			//这里就在 {{.title}} 模板中使用
			"title" : "users/index",
		})
	})

	//启动HTTP服务，默认地址是127.0.0.1:9090启动服务
	r.Run(":8080")

}


//目前的页面都很丑，可以在网上直接下载一个前端模板，就可以很炫酷了。








