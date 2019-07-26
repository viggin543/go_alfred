package main

//go:generate protoc3 -I banana --go_out=plugins=grpc:banana banana/banana.proto

func main() {
	grpcServerExample()
	httpServerExample()
}
