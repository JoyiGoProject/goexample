package client

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"

	"github.com/astaxie/beego"
	_ "github.com/hprose/hprose-golang/rpc/websocket"
)

type ErrorEvent struct{}

func (e *ErrorEvent) OnError(name string, err error) {
	fmt.Println(name, err.Error())
}

type TestFunc struct {
	//检查校验码  servercode 校验码
	CheckServerCode func(servercode string) (b bool)
	//检查连接状态
	CheckServerStatus func() error
	//其它方法
	//...
}

func Test(serverip, serverport string) (clientTest *TestFunc, err error) {
	url := "ws://" + serverip + ":" + serverport + "/rpc"
	beego.Debug("request url", url)
	client := rpc.NewClient(url)
	client.SetEvent(ErrorEvent{})
	client.UseService(&clientTest)
	err = clientTest.CheckServerStatus()
	return clientTest, err
}
