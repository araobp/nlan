package main

import (
	"bytes"
	"flag"
	"log"
	_ "os"

	nlan "github.com/araobp/golan/nlan/model/nlan"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":8282"
)

func main() {
	target := flag.String("target", "localhost", "target host")
	flag.Parse()
	var buffer bytes.Buffer
	buffer.WriteString(*target)
	buffer.WriteString(port)
	address := buffer.String()
	log.Printf("target: %s\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Print(err)
	}
	agent := nlan.NewNlanAgentClient(conn)

	dvr := nlan.Dvr{OvsBridges: true}
	model := nlan.Model{Dvr: &dvr}
	request := nlan.Request{Model: &model}
	response, err := agent.Add(context.Background(), &request)
	if err != nil {
		log.Print(err)
	}
	log.Print(response)

	ptn := nlan.Ptn{}
	model = nlan.Model{Ptn: &ptn}
	request = nlan.Request{Model: &model}
	response, err = agent.Add(context.Background(), &request)
	if err != nil {
		log.Print(err)
	}
	log.Print(response)
}
