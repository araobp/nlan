// Master
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/araobp/go-nlan/nlan/env"
	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
	st "github.com/araobp/go-nlan/nlan/state"
	"github.com/araobp/go-nlan/nlan/util"
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
	cont := context.Background()

	// Hello
	cs := []string{"hello"}
	cs_master := nlan.Capabilities{Capability: cs}
	cs_agent, err := agent.Hello(cont, &cs_master)
	if err != nil {
		log.Print(err)
	}
	log.Println(cs_agent)

	// NLAN Request
	request := nlan.Request{Model: model}
	var response *nlan.Response
	switch ope {
	case env.ADD:
		response, err = agent.Add(cont, &request)
	case env.UPDATE:
		response, err = agent.Update(cont, &request)
	case env.DELETE:
		response, err = agent.Delete(cont, &request)
	}
	if err != nil {
		log.Print(err)
	}
	log.Print(response)
	r := result{address: address, ope: ope, state: model, response: response}
	c <- r
}

func main() {
	filename := flag.String("state", "state.yaml", "state file")
	service := flag.String("service", "ptn", "model")
	flag.Parse()
	log.Println(*filename)
	log.Println(*service)

	// Reads the state file
	state, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(state))

	// Converts YAML to Go struct
	var hosts map[string]interface{} = util.ListHosts()
	fmt.Println(hosts)
	state_ := st.NetworkState{}
	statestring := string(state)
	util.Yaml2Struct(&statestring, &state_, hosts)
	fmt.Println(state_)

	// Deployment
	for _, v := range state_.States {
		router := v.Router
		ip := string(hosts[router].(string))
		var buffer bytes.Buffer
		buffer.WriteString(ip)
		buffer.WriteString(env.PORT)
		address := buffer.String()
		model := v.Model
		c := make(chan result)
		go deploy(address, env.ADD, &model, c)
		r, _ := <-c
		r.Println()
	}
}
