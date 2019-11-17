package main

import (
	"context"
	"fmt"
	"github.com/dollarkillerx/GoMicro-Study/hellogrpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
 * this is grpc hello demo server
 */

type HelloService struct {

}

func (h *HelloService) SyaHello(ctx context.Context,request *pb.HelloRequest) (*pb.HelloResponse,error) {
	return &pb.HelloResponse{
		Name:"this is test by grpc!",
	}, nil
}

func main() {
	host := fmt.Sprintf(":%d", 9000)
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server,&HelloService{})

	log.Println("server run is: ",host)
	if err := server.Serve(listen);err != nil {
		log.Fatalln(err)
	}

}