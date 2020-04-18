package main

import (
	"fmt"
	_ "github.com/cihub/seelog"
	"os"
)

func main() {
	// “Info”只是Seelog支持的日志级别之一。
	//还可以使用Trace, Debug, Info, Warn, Error, Critical levels.
	//log.Info("hello info")
	//log.Trace("hello trace")
	//log.Debug("hello debug")
	//log.Warn("hello warn")
	//log.Error("hello error")
	//log.Critical("hello Critical")
	//
	////加载配置信息
	//logger, err := log.LoggerFromConfigAsBytes()
	//if err != nil {
	//	return err
	//}
	//log.ReplaceLogger(logger)
	//defer log.Flush()

	fmt.Println(os.Args[0])
}











