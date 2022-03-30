package main

import (
	"fmt"
	"github.com/channingduan/rpc-skeleton/models"
	"github.com/channingduan/rpc-skeleton/service"
	"github.com/channingduan/rpc/config"
	"github.com/channingduan/rpc/database"
	"github.com/channingduan/rpc/server"
	"github.com/oscto/ky3k"
	"github.com/smallnest/rpcx/log"
	"os"
)

func main() {

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	conf, err := config.Register("./server.json")
	fmt.Println(ky3k.JsonToString(conf))
	if err != nil {
		log.Errorf("config parse error: %v", conf)
	}

	var srvJerry config.Method
	srvJerry.Name = "jerry"
	srvJerry.Func = service.Jerry

	var srvHuman config.Method
	var f service.Human
	srvHuman.Name = "hello.world"
	srvHuman.Func = f.Hello

	db := database.Register(conf)
	err = db.AutoMigrate(&models.User{}, &models.ShippingAddress{})
	if err != nil {
		fmt.Println("AutoMigrate: ", err)
		os.Exit(1)
	}
	srv := server.NewServer(conf)

	srv.AddMethod(srvJerry)
	srv.AddMethod(srvHuman)
	srv.Start()
}
