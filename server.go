package main

import (
	"github.com/channingduan/rpc/config"
	"github.com/channingduan/rpc/server"
	"github.com/channingduan/rpc/service"
)

func main() {

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	conf := config.Config{
		BasePath:     "rpc",
		ServicePath:  "service",
		ServiceName:  "test",
		ServiceAddr:  "127.0.0.1:8089",
		RegistryAddr: "127.0.0.1:8500",
	}

	var srvJerry config.Method
	srvJerry.Name = "jerry"
	srvJerry.Func = service.Jerry

	var srvHuman config.Method
	var f service.Human
	srvHuman.Name = "hello.world"
	srvHuman.Func = f.Hello

	srv := server.NewServer(&conf)
	srv.AddMethod(srvJerry)
	srv.AddMethod(srvHuman)
	srv.Start()
}
