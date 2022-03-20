package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "hello/proto/hello"
)

const Address  = "127.0.0.1:50052"

func main()  {
	//连接
	conn,err:=grpc.Dial(Address,grpc.WithInsecure())
	if err!=nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	
	//初始化客户端
	c:=pb.NewHelloClient(conn)
	
	//调用方法
	req:=&pb.HelloRequest{Name:"gRPC"}
	res,err:=c.SayHello(context.Background(),req)
	if err!=nil {
		grpclog.Fatalln(err)
	}
	fmt.Print(res.Message)
	grpclog.Println(res.Message)
}

