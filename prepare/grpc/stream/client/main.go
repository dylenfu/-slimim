package main

import (
	"flag"
	"io"
	"log"
	"strconv"

	pb "slimim/prepare/grpc/stream/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String("endpoint", "localhost:9091", "grpc server address")
	testCase = flag.Int("switch", 1, "chose test case")
)

func main() {
	flag.Parse()

	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(*endpoint, opt)
	if err != nil {
		log.Fatalf("failed to dail %v", err)
	}
	defer conn.Close()

	client := pb.NewApiServiceClient(conn)
	switch *testCase {
	case 1:
		hello1(client)
	case 2:
		hello2(client)
	case 3:
		hello3(client)
	}
}

func hello1(c pb.ApiServiceClient) {
	stream, err := c.SayHello1(context.Background(), &pb.HelloReq{Data: "first"})
	if err != nil {
		log.Fatalf("failed to greet %v", err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to recv %v", err)
		}

		log.Printf("Greeting is %s", reply.Data)
	}
}

func hello2(c pb.ApiServiceClient) {
	stream, err := c.SayHello2(context.Background())
	if err != nil {
		log.Fatalf("failed to greet %v", err)
	}

	for i := 0; i < 100; i++ {
		stream.Send(&pb.HelloReq{Data: "second"})
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recv %v", err)
	}

	log.Printf("greeting: %s", reply.Data)
}

func hello3(c pb.ApiServiceClient) {
	stream, err := c.SayHello3(context.Background())
	if err != nil {
		log.Fatalf("failed to call %v", err)
		return
	}

	var i int
	for {
		stream.Send(&pb.HelloReq{Data: "third " + strconv.Itoa(i)})
		reply, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to recv %v", err)
			break
		}

		log.Fatalf("Greeting %s", reply.Data)
		i++
	}
}
