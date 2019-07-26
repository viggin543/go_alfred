package main

import (
	"context"
	pb "example.com/banana/server/main/banana"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

const port = ":50051"

func sendGrpcRequestToSErver() {
	conn, _ := grpc.Dial("localhost"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewBananaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bananaReply, _ := client.Greet(ctx, &pb.BananaRequest{Name: "yellow", Size: 123})
	fmt.Println("GRPC IS FUN",bananaReply)
}

func main(){
	sendGrpcRequestToSErver()
}