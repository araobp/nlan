// Master
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/araobp/nlan/common"
	"github.com/araobp/nlan/env"
	nlan "github.com/araobp/nlan/model/nlan"
	"github.com/araobp/nlan/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type result struct {
	router   string
	address  string
	ope      int
	state    *nlan.Model
	response *nlan.Response
}

func (r *result) Println() {
	fmt.Println("---")
	var state nlan.Model
	switch r.ope {
	case env.CLEAR:
		state = nlan.Model{} // empty
	default:
		state = *r.state
	}
	response := *r.response
	exit := response.Exit
	log := response.LogMessage
	fmt.Printf("router: %s, address: %s\nope: %d\nstate: %s\nexit: %d\n", r.router, r.address, r.ope, state, exit)
	fmt.Println(log)
}

var channel = make(chan result)
var wg sync.WaitGroup
var count int

func deploy(router string, address string, ope int, state *nlan.Model) {
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
	request := nlan.Request{Model: state}
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
	r := result{router: router, address: address, ope: ope, state: state, response: response}
	channel <- r
}

func clearConfig(router string, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	agent := nlan.NewNlanAgentClient(conn)
	cont := context.Background()

	// Clear
	clearMode := nlan.ClearMode{Terminate: false}
	response, err := agent.Clear(cont, &clearMode)
	if err != nil {
		log.Print(err)
	}
	log.Println(response)

	r := result{router: router, address: address, ope: env.CLEAR, response: response}

	channel <- r
}

func main() {
	count = 0
	filename := flag.String("state", "state.yaml", "state file")
	clear := flag.Bool("clear", false, "clear config at NLAN agent")
	reset := flag.Bool("reset", false, "reset NLAN state on etcd")
	flag.Parse()

	if *reset == true {
		util.ResetState()
		os.Exit(0)
	}

	log.Println(*filename)

	states, hosts := common.ReadState(filename, nil)

	// Deployment
	for _, v := range *states {
		router := v.Router
		hosts_ := *hosts
		ip := string(hosts_[router].(string))
		var buffer bytes.Buffer
		buffer.WriteString(ip)
		buffer.WriteString(env.PORT)
		address := buffer.String()
		state := v.Model
		switch *clear {
		case true:
			count += 1
			go clearConfig(router, address)
		case false:
			count += 1
			util.SetMode(router, env.INIT)
			go deploy(router, address, env.ADD, &state)
		}
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("---------------------------")
		for r := range channel {
			r.Println()
			util.SetState(r.router, r.state)
			util.SetMode(r.router, env.RESTART)
			count -= 1
			if count <= 0 {
				break
			}
		}
	}()
	wg.Wait()
}
