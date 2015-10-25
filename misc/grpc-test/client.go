package main

import (
	"log"
	_ "os"

	api "github.com/araobp/go-nlan/misc/grpc-test/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8282"
)

func main() {

	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	c := api.NewNlanAgentClient(conn)

	bridges := api.Request_Bridges{Bridges: &api.Bridges{Dummy: "test"}}
	request := api.Request{State: &bridges}
	reply, _ := c.Add(context.Background(), &request)
	log.Printf("%d\n", reply.Result)
	log.Printf("%s\n", reply.LogMessage)
}
