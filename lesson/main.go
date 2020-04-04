// 代码主要讲了Go的模板，使用模板去传递一个变量、结构体、map; 自定义函数、模板嵌套
// 另外对于模板的愈发的介绍就主要看李文周的博客：https://www.liwenzhou.com/posts/Go/go_template/

package main

import (
	"fmt"
	"html/template"
	"net/http"
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

//传递一个 struct 和 map
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

//模板中自定义函数
func saymyfunc(w http.ResponseWriter, r *http.Request) {
	//定义一个函数k
	//		要么只有一个返回值，要么有两个返回值，第二个返回值只能是error类型
	k := func(name string) (string,error) {
		return name + "好看", nil 
	}

	//定义模板
	t := template.New("f.tmpl")

	//告诉模板引擎，现在多了一个自定义的函数k
	t.Funcs(template.FuncMap{
		"kua99" : k,
	})

	//解析模板
	//		创建一个名为f.tmpl的模板对象
	//		New中的文件名至少要和ParseFiles中的一个文件名对应上
	_,err := t.ParseFiles("./templates/posts/f.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles failed,err:%v",err)
		return
	}
	//渲染模板
	name := "haha"
	t.Execute(w,name)
}

//模板嵌套
func demo(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	//		这里文件的顺序一定是被包含的写在后面，先写妈妈
	t,err := template.ParseFiles("./templates/posts/g.tmpl","./templates/posts/ul.tmpl")
	if err != nil {
		fmt.Println("failed")
		return
	}
	name := "zcy"
	//渲染模板
	t.Execute(w,name)
}


func main() {
	http.HandleFunc("/",sayhello)
	http.HandleFunc("/struct",saystruct)
	http.HandleFunc("/map",saymap)
	http.HandleFunc("/structandmap",saystructandmap)

	//自定义函数
	http.HandleFunc("/myfunc",saymyfunc)
	//模板嵌套
	http.HandleFunc("/tmplDemo",demo)




	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("HTTP failed")
	}
}


















