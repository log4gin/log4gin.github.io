package main

import (
	"context"
	"fmt"
	"grpc_hello/protos/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedGreeterServer //兼容之前的版本
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloReq) (*hello.HelloResp, error) {
	fmt.Println("get msg :", in.Name)
	return &hello.HelloResp{
		Resp: "Hello " + in.Name + ", i am go grpc server",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":2333")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	log.Fatal(s.Serve(listen))
}
