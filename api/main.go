package main

import (
	"context"
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

type podServiceClient struct {
	pb.UnimplementedPodServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHealthServiceServer(s, &healthServiceClient{})
	pb.RegisterPodServiceServer(s, &podServiceClient{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *healthServiceClient) GetStatus(ctx context.Context, in *pb.Empty) (*pb.AliveResponse, error) {
	return &pb.AliveResponse{Status: true}, nil
}
