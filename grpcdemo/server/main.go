package main


import (
	"fmt"
	pb "grpcdemo/pb"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReplay, error) {
	return &pb.HelloReplay{Message:"Hello" + request.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Println("failed to listen : %v", err)
		return
	}
	//创建RPC服务器
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})	//在rpc服务端注册服务

	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve:%v",err)
		return
	}

}

