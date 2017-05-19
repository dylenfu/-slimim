package main

import (
	"log"
	"net"
	pb "slimim/prepare/grpc/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9090"
	cert = "../resources/cert.pem"
	key  = "../resources/key.pem"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("failed to server tls from file %v", err)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)

	pb.RegisterApiServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{"hello" + request.Data}, nil
}
