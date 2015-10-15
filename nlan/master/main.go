// Master
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	_ "os"

	env "github.com/araobp/golan/nlan/env"
	nlan "github.com/araobp/golan/nlan/model/nlan"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func deploy(address string, ope int, state *string, c chan<- int) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	agent := nlan.NewNlanAgentClient(conn)

	// Reads the state file and converts it into a Go struct
	json_data, err := ioutil.ReadFile(*state)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(json_data))
	//var model nlan.Model
	model := nlan.Model{}
	if err := json.Unmarshal(json_data, &model); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", model)

	// NLAN Request
	request := nlan.Request{Model: &model}
	var response *nlan.Response
	switch ope {
	case env.ADD:
		response, err = agent.Add(context.Background(), &request)
	case env.UPDATE:
		response, err = agent.Update(context.Background(), &request)
	case env.DELETE:
		response, err = agent.Delete(context.Background(), &request)
	}
	if err != nil {
		log.Print(err)
	}
	log.Print(response)
	c <- 0
}

func main() {
	target := flag.String("target", "localhost", "target host")
	state := flag.String("state", "state.json", "state file")
	service := flag.String("service", "ptn", "model")
	flag.Parse()
	log.Println(*target)
	log.Println(*state)
	log.Println(*service)

	// Connects to the target host
	var buffer bytes.Buffer
	buffer.WriteString(*target)
	buffer.WriteString(env.PORT)
	address := buffer.String()
	log.Printf("target: %s\n", address)

	// Deployment
	c := make(chan int)
	go deploy(address, env.ADD, state, c)
	value, _ := <-c
	fmt.Printf("chan: %d\n", value)
}
