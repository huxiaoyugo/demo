package main

import (
	"context"
	micro2 "demo/micro"
	"fmt"
	"github.com/micro/go-micro"
	"log"
)


type tp struct {
}


func (s tp) SayHello(ctx context.Context, req *micro2.HelloReq, res *micro2.HelloRes) error {
	str := fmt.Sprintf("hello %s, %s", req.Name, req.Content)
	res.Status = micro2.RetStatus_RetStatus_Success
	res.Content = str
	log.Print(str)
	return nil
}

func (s tp) SayBye(ctx context.Context, req *micro2.ByeReq, res *micro2.ByeRes) error {
	log.Print(req.Content)
	str := fmt.Sprintf("bye-bye %s", req.Name)
	res.Status = micro2.RetStatus_RetStatus_Success
	res.Content = str
	log.Print(str)
	return nil
}


func main() {


	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉

	//reg := consul.NewRegistry(func(op *registry.Options){
	//	op.Addrs = []string{
	//		"http://localhost:2379",
	//	}
	//})


	srv := micro.NewService(
		micro.Name("pb"),
		micro.Version("latest"),
		//micro.Registry(reg),
	)
	srv.Init()
	t := &tp{}
	micro2.RegisterTpHandler(srv.Server(), t)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}