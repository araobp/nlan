package main

import (
	"log"
	"net"

	nlan "github.com/araobp/golan/nlan/model/nlan"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":8282" // NLAN agent listen port
)

type agent struct{}

// Add method
func (a *agent) Add(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	log.Print("Server: Add() is called")
	response := nlan.Response{Result: 0, LogMessage: "Server: Add() is called"}
	return &response, nil
}

// Update method
func (a *agent) Update(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	log.Print("Server: Update() is called")
	response := nlan.Response{Result: 0, LogMessage: "Server: Add() is called"}
	return &response, nil
}

// Delete method
func (a *agent) Delete(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	log.Print("Server: Delete() is called")
	response := nlan.Response{Result: 0, LogMessage: "Server: Add() is called"}
	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Print(err)
	}
	s := grpc.NewServer()
	nlan.RegisterNlanAgentServer(s, &agent{})
	s.Serve(listen)
}
