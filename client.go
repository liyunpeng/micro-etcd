package main

import (
	"context"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro"
	model "proto"
	"fmt"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			// 经试验, 不用配置这些地址， 只要本地启动etcd, micro的业务提供者和使用者都可正常工作
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()

	sayClent := model.NewSayService("lp.srv.eg1", service.Client())

	rsp, err := sayClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}
