package main

import (
	"context"
	pb "example.com/banana/server/main/banana"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

func (s *server) Greet(ctx context.Context, in *pb.BananaRequest) (*pb.BananaReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.BananaReply{Message: "Banana " + in.Name}, nil
}

func grpcServerExample() {
	createServer()
	sendGrpcRequestToSErver()

}

const port = ":50051"

func sendGrpcRequestToSErver() {
	conn, _ := grpc.Dial("localhost"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewBananaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bananaReply, _ := client.Greet(ctx, &pb.BananaRequest{Name: "yellow", Size: 123})
	fmt.Println("GRPC IS FUN", bananaReply)
}

func createServer() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterBananaServer(s, &server{})
	fmt.Println("server started")
	_ = s.Serve(lis)

}
