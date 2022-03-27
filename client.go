package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/channingduan/gateway/service"
	"github.com/channingduan/rpc/client"
	"github.com/channingduan/rpc/config"
	rpcClient "github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/codec"
	"github.com/smallnest/rpcx/protocol"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//callRpc()
	callGateway()
}
func call() {

	discovery, err := rpcClient.NewConsulDiscovery(config.BasePath, "service", []string{"127.0.0.1:8500"}, nil)
	if err != nil {
		fmt.Println("NewDiscovery err: ", err)
	}

	opt := rpcClient.DefaultOption
	opt.SerializeType = protocol.JSON
	rClient := rpcClient.NewXClient("service", rpcClient.Failtry, rpcClient.RandomSelect, discovery, opt)
	defer rClient.Close()

	args := config.Request{
		Message: "Hello world",
	}
	reply := &config.Response{}
	if err = rClient.Call(context.Background(), "jerry", args, reply); err != nil {
		log.Fatalf("rpc call err: ", err)
	}

	log.Printf("request: %v => response: %v", args, reply)
}

func callRpc() {
	args := config.Request{
		Message: "Hello world",
	}
	conf := config.Config{
		BasePath:     "rpc",
		ServicePath:  "service",
		ServiceName:  "test",
		ServiceAddr:  "127.0.0.1:8089",
		RegistryAddr: "127.0.0.1:8500",
	}

	xClient := client.NewClient(&conf)
	reply, err := xClient.Call(context.Background(), conf.ServicePath, "hello", args)
	if err != nil {
		fmt.Println("Call error: ", err)
	}
	log.Printf("request: %v => response: %v", args, reply)

	reply, err = xClient.Call(context.Background(), conf.ServicePath, "jerry", args)
	if err != nil {
		fmt.Println("Call error: ", err)
	}
	log.Printf("request: %v => response: %v", args, reply)
}

func callGateway() {

	cc := &codec.MsgpackCodec{}
	args := config.Request{
		Message: "Hello world",
	}
	data, _ := cc.Encode(args)

	var serviceName = "service"
	var serviceMethod = "hello"

	url := fmt.Sprintf("http://127.0.0.1:9090/%s/%s", "hello", "world")
	fmt.Println("url", url)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed to create request: ", err)
		return
	}

	h := req.Header
	// h.Set(gateway.XMessageID, strconv.Itoa(rand.Int()))
	// h.Set(gateway.XMessageType, "0")
	h.Set(service.XSerializeType, "3")
	h.Set(service.XServicePath, serviceName)
	h.Set(service.XServiceMethod, serviceMethod)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("#%d failed to call: ", err)
	}
	defer res.Body.Close()

	// handle http response
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	reply := &config.Response{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Println("请求正常: ", reply.Message)
}
