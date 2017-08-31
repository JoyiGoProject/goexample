package main

import (
	"goexample/src/task"
	"os"
	"tlog"

	"github.com/astaxie/beego"
)

func main() {
	//需要在启动时执行的方法,这是一个并发的方法，加上关键字go
	go task.Init()
	initLog()
	//程序启动
	beego.Run()
}

//tlog日志配置
func initLog() {
	if _, err := os.Stat("log"); err != nil {
		os.Mkdir("log", 0755)
	}
	tlog.SetLogger("file", `{"filename":"log/hcloud.log"}`)
}
