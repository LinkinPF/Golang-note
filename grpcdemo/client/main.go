package main

import (
	"context"
	"fmt"

	pb "grpcdemo/pb"
	"google.golang.org/grpc"
)

func main() {
	//连接服务器
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed tp connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	// 调用服务端的SayHello方法
	r,err := c.SayHello(context.Background(), &pb.HelloRequest{Name:"zcy"})
	if err != nil {
		fmt.Printf("could not greet : %v", err)
	}

	fmt.Printf("Greeting: %s !\n", r.Message)


}
