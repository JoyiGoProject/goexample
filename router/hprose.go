package router

import (
	"goexample/api"
	"rpc/server"

	"github.com/astaxie/beego"
)

func init() {
	startServer()
}

//导出对外的api
func startServer() {
	server.Server.AddFunction("UserList", api.UserList)

	beego.Handler("/rpc", server.Server)
}
