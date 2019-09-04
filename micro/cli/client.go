package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"

	micro2 "demo/micro"
)

//import "demo/micro"

func main() {
	srv := micro.NewService(
		micro.Name("pb"),
		micro.Version("latest"),
	)
	srv.Init()
	cli := micro2.NewTpClient("pb", srv.Client())


	req := &micro2.HelloReq{
		Name:                 "胡小于",
		Content:              "再见啦",
	}
	res, _ := cli.SayHello(context.Background(), req)
	fmt.Printf("%s",res.Content)
}
