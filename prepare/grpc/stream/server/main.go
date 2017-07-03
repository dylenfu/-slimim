package main

import (
	"flag"
	"io"
	"log"
	"net"

	pb "slimim/prepare/grpc/stream/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.String("port", ":9091", "grpc server listen port")
)

type server struct{}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	defer s.Stop()

	pb.RegisterApiServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (s *server) SayHello1(request *pb.HelloReq, gs pb.ApiService_SayHello1Server) error {
	for i := 0; i < 100; i++ {
		gs.Send(&pb.HelloResp{Data: "Hello1 " + request.Data})
	}
	return nil
}

func (s *server) SayHello2(gs pb.ApiService_SayHello2Server) error {
	var names []string
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			gs.SendAndClose(&pb.HelloResp{Data: "io eof"})
			return nil
		}
		if err != nil {
			log.Fatalf("failed to Recv %v", err)
			return err
		}
		names = append(names, "Hello2 "+in.Data)
	}
	return nil
}

func (s *server) SayHello3(gs pb.ApiService_SayHello3Server) error {
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to Recv %v", err)
			return err
		}
		gs.Send(&pb.HelloResp{Data: "Hello3 " + in.Data})
	}
	return nil
}
