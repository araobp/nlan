package main

import (
	"log"
	"net"

	hw "github.com/araobp/gone/helloworld/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":8282"
)

type server struct{}

func (s *server) HelloWorld(ctx context.Context, in *hw.HelloRequest) (*hw.HelloReply, error) {
	return &hw.HelloReply{Name: in.Name, Greeting: in.Greeting}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	hw.RegisterHelloWorldServiceServer(s, &server{})
	s.Serve(listen)
}
