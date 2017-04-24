package main

import (
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc"
	pb "slimim/prepare/grpc/rpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9090"

type server struct{}

func Server() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApiServiceServer(s, server{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context,request *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{"hello" + request.Data}, nil
}