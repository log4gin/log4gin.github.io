# Go下的RPC的基本使用

### 安装必要工具

1. protobuf 用于编译proto文件,直接在ubuntu的应用商店安装

2. ```sh
   # protoc-gen-go 用于序列化和反序列化
   go install google.golang.org/protobuf/cmd/protoc-gen-go
   # protoc-gen-go-grpc 用于生成客户端和服务端的代码
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
   ```

### 编写hello.proto文件

```protobuf
// 指定版本
syntax = "proto3";

// 输出包路径
option go_package = "./;hello";

message HelloReq {
    string Name = 1;   // 1 为标识id
}

message HelloResp {
    string Resp = 1;
}

service Greeter {
    rpc SayHello (HelloReq) returns (HelloResp) {}
}
```

### 编译

```
proto ./hello.proto --go_out=.
proto ./hello.proto --go-grpc_out=.
```

### 编写服务端

```
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
		Resp: "Hello " + in.Name,
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

```

### 编写客户端

```
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

	rsp, err := c.SayHello(ctx, &hello.HelloReq{Name: "log4gin"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.GetResp())
}

```

[过程源码](https://github.com/log4gin/log4gin.github.io/tree/main/src/go_grpc_hello)