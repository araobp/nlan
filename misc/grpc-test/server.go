package main

import (
	"log"
	"net"

	api "github.com/araobp/go-nlan/misc/grpc-test/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":8282"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *api.Request) (*api.Reply, error) {
	log.Print("Server: Add() is called")
	bridges := in.GetBridges()
	log.Printf("Dummy: %s\n", bridges.Dummy)
	reply := api.Reply{Result: 0, LogMessage: "Server: Add() is called"}
	return &reply, nil
}

func (s *server) Update(ctx context.Context, in *api.Request) (*api.Reply, error) {
	log.Print("Server: Update() is called")
	reply := api.Reply{Result: 0, LogMessage: "Server: Update() is called"}
	return &reply, nil
}

func (s *server) Delete(ctx context.Context, in *api.Request) (*api.Reply, error) {
	log.Print("Server: Delete() is called")
	reply := api.Reply{Result: 0, LogMessage: "Server: Delete() is called"}
	return &reply, nil
}

func main() {
	listen, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	api.RegisterNlanAgentServer(s, &server{})
	s.Serve(listen)
}
