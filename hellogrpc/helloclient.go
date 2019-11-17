package main

import (
	"context"
	"github.com/dollarkillerx/GoMicro-Study/hellogrpc/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

/**
 * this is hello server client
 */

func main() {
	host := ":9000"
	conn,err := grpc.Dial(host,grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	for {
		response, err := client.SyaHello(context.TODO(), &pb.HelloRequest{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(response.Name)
		time.Sleep(time.Millisecond * 200)
	}
}