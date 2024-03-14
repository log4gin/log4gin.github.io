package main

import (
	"context"
	"fmt"
	"grpc_hello/protos/hello"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("localhost:2333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer con.Close()
	c := hello.NewGreeterClient(con)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rsp, err := c.SayHello(ctx, &hello.HelloReq{Name: "tomygin"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.GetResp())
}
