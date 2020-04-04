// 代码主要讲了Go的模板，使用模板去传递一个变量、结构体、map、
// 另外对于模板的愈发的介绍就主要看李文周的博客：https://www.liwenzhou.com/posts/Go/go_template/

package main

import (
	"fmt"

	_ "github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

type student struct {
	Name string
	Age int
	Tel int
}

// 传递一个字符串
func sayhello(w http.ResponseWriter, r *http.Request) {
	t,err := template.ParseFiles("./templates/posts/index.tmpl")
	if err != nil {
		fmt.Println("failed")
	}
	err = t.Execute(w,"hh")
	if err != nil {
		fmt.Println("failed")
		return
	}
}

// 传递一个结构体
func saystruct(w http.ResponseWriter, r *http.Request) {
	t,err := template.ParseFiles("./templates/posts/struct.tmpl")
	if err != nil {
		fmt.Println("1 failed")
	}
	//这里注意，如果字段名称是小写的，那么在网页上就不会显示了。
	//而且如果发送私有的字段，t.Execute还会报错的
	user1 := student{
		Name : "zcy",
		Age : 10,
		Tel : 15686078926,
	}

	err = t.Execute(w,user1)
	if err != nil {
		fmt.Println("2 failed")
		return
	}
}

// 传递一个map
func saymap(w http.ResponseWriter, r *http.Request) {
	t,err := template.ParseFiles("./templates/posts/map.tmpl")
	if err != nil {
		fmt.Println("1 failed")
	}
	
	//使用map的时候一半就用小写字母了。
	map1 := map[string]interface{} {
		"name" : "zcy",
		"age" : 18,
		"gender" : "male",
	}

	err = t.Execute(w,map1)
	if err != nil {
		fmt.Println("2 failed")
		return
	}
}

func saystructandmap(w http.ResponseWriter, r *http.Request) {
	t,err := template.ParseFiles("./templates/posts/structandmap.tmpl")
	if err != nil {
		fmt.Println("1 failed")
	}
	user1 := student{
		Name : "zcy",
		Age : 10,
		Tel : 15686078926,
	}
	map1 := map[string]interface{} {
		"name" : "zcy",
		"age" : 18,
		"gender" : "male",
	}
	err = t.Execute(w,map[string]interface{} {
		"u1" : user1,
		"map1" : map1,
	})
	if err != nil {
		fmt.Println("2 failed")
		return
	}
}

func main() {
	http.HandleFunc("/",sayhello)
	http.HandleFunc("/struct",saystruct)
	http.HandleFunc("/map",saymap)
	http.HandleFunc("/structandmap",saystructandmap)
	

	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("HTTP failed")
	}
}























