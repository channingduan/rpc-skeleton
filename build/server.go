package main

import (
	"github.com/channingduan/rpc-skeleton/controller"
	"github.com/channingduan/rpc-skeleton/models"
	"github.com/channingduan/rpc-skeleton/router"
	"github.com/channingduan/rpc/config"
	"github.com/channingduan/rpc/database"
	"github.com/channingduan/rpc/server"
	"github.com/smallnest/rpcx/log"
)

type Router struct {
	Name  string `json:"name"`
	Roter string
}

func main() {

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	// 初始化配置文件
	conf, err := config.Register("./server.json")
	if err != nil {
		log.Errorf("config parse error: %v", conf)
	}
	// 初始化数据库
	db := database.Register(conf)
	_ = db.AutoMigrate(&models.User{}, &models.ShippingAddress{})
	srv := server.NewServer(conf)
	router.Initial(srv, controller.Register(conf))
	srv.Start()
}
