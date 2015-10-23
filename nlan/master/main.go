// Master
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/araobp/golan/nlan/env"
	nlan "github.com/araobp/golan/nlan/model/nlan"
	"github.com/araobp/golan/nlan/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type result struct {
	address  string
	ope      int
	state    *nlan.Model
	response *nlan.Response
}

func (r *result) Println() {
	fmt.Println("---")
	address := r.address
	ope := r.ope
	state := *r.state
	response := *r.response
	exit := response.Exit
	log := response.LogMessage
	fmt.Printf("address: %s\nope: %d\nstate: %s\nexit: %d\nlog: %s\n", address, ope, state, exit, log)
}

func deploy(address string, ope int, model *nlan.Model, c chan<- result) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	agent := nlan.NewNlanAgentClient(conn)

	// NLAN Request
	request := nlan.Request{Model: model}
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
	r := result{address: address, ope: ope, state: model, response: response}
	c <- r
}

func main() {
	target := flag.String("target", "localhost", "target host")
	filename := flag.String("state", "state.json", "state file")
	service := flag.String("service", "ptn", "model")
	flag.Parse()
	log.Println(*target)
	log.Println(*filename)
	log.Println(*service)

	// Connects to the target host
	var buffer bytes.Buffer
	buffer.WriteString(*target)
	buffer.WriteString(env.PORT)
	address := buffer.String()
	log.Printf("target: %s\n", address)

	// Reads the state file
	state, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(state))

	// Converts YAML to Go struct
	var hosts map[string]interface{} = util.ListHosts()
	fmt.Println(hosts)
	model := nlan.Model{}
	statestring := string(state)
	util.Yaml2Struct(&statestring, &model, hosts)

	// Deployment
	c := make(chan result)
	go deploy(address, env.ADD, &model, c)
	r, _ := <-c
	r.Println()
}
