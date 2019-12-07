package main

import (
	"context"
	"github.com/micro/go-plugins/registry/etcdv3"
	model "proto"
	"fmt"
)

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	// 初始化服务
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
