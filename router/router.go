package router

import (
	"github.com/channingduan/rpc-skeleton/controller"
	"github.com/channingduan/rpc/config"
	"github.com/channingduan/rpc/server"
)

func Initial(rpcServer *server.RpcServer, ctx *controller.Controller) {

	rpcServer.AddMethod(config.Method{
		Name:   "用户登录",
		Router: "user.login",
		Func:   ctx.Login,
	})
	rpcServer.AddMethod(config.Method{
		Name:   "用户注册",
		Router: "user.register",
		Func:   ctx.Register,
	})
	rpcServer.AddMethod(config.Method{
		Name:   "个人信息",
		Router: "user.profile",
		Func:   ctx.Profile,
	})
}
