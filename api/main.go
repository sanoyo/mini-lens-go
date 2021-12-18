package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sanoyo/mini-lens-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type healthServiceClient struct {
	pb.UnimplementedHealthServiceServer
}

func NewServer() *healthServiceClient {
	return &healthServiceClient{}
}

func (s *healthServiceClient) GetStatus(ctx context.Context, in *pb.Empty) (*pb.AliveResponse, error) {
	return &pb.AliveResponse{Status: true}, nil
}

func main() {
	fmt.Println("Listen Address:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHealthServiceServer(s, &healthServiceClient{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
