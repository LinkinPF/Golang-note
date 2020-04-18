package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)


func main() {
	app := cli.NewApp()
	//下面两个参数会在-h命令选项中体现
	app.Name = "lmp"
	app.Usage = "lmp_test"

	//Flags就是设置命令行的参数了,不同类型的参数主要体现在Value值得类型不同上
	app.Flags = []cli.Flag {
		//使用 -h 查看help，下面这个就会显示：
		//--lang value, -l value  language for the greeting (default: "english")
		cli.StringFlag{
			Name:        "lang,l",
			Usage:       "language for the greeting",
			Value:       "english",
		},
	}

	//Action函数就是程序具体的执行
	app.Action = func(c *cli.Context) error {
		name := "zcy"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		//Context有很多的方法，c.String(name)就是在输入的命令行参数中查找名为name，类型为StringFlag的参数
		//类似于这样的还有Uint，查找UintFlag类型的参数
		if c.String("lang") == "english" {
			fmt.Println("hello",name)
		} else {
			fmt.Println("你好",name)
		}
		//使用这种方法就可以获得命令行写入的内容，不过这里不能写-开头的
		fmt.Printf("args:%q\n",c.Args().Get(0))
		return nil
	}



	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
