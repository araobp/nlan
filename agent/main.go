package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	config_ptn "github.com/araobp/nlan/agent/config/ptn"
	config_router "github.com/araobp/nlan/agent/config/router"
	config_vhosts "github.com/araobp/nlan/agent/config/vhosts"
	con "github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/common"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"
	st "github.com/araobp/nlan/state"
	"github.com/araobp/nlan/util"
	"golang.org/x/net/context"
)

type agent struct {
	con *con.Context
}

func (a *agent) route(crud int, in *nlan.Request, configMode int) (exit uint32) {
	model := in.Model
	ptn := model.GetPtn()
	dvr := model.GetDvr()
	vhosts := model.GetVhosts()
	router := model.GetRouter()

	exit = 0
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			exit = 1
		}
	}()

	if ptn != nil {
		log.Print("Routing to PTN module...")
		config_ptn.Crud(crud, ptn, a.con)
	}
	if dvr != nil {
		//
	}
	if vhosts != nil {
		log.Print("Routing to VHOSTS module...")
		config_vhosts.Crud(crud, vhosts, a.con)
	}
	if router != nil {
		log.Print("Routing to ROUTER module...")
		config_router.Crud(crud, router, a.con)
	}
	return exit
}

func main() {

	// Command options
	target := os.Getenv("HOSTNAME")
	ope := flag.String("ope", "ADD", "CRUD operation")
	filename := flag.String("state", "", "state file")
	flag.Parse()

	// Log file setup
	file := fmt.Sprintf("/var/volume/nlan-agent-%s.log", target)
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Unable to open log file: %s", file)
	}
	log.SetOutput(f)

	// Registers host IP address with tega db
	router := util.RegisterHost()

	// Gets start mode from teag db
	log.Printf("Start mode: %d", mode)
	cmd, cmdp := util.GetCmd(true)

	//Adds a secondary IP address to eth0
	secondary := util.GetSecondaryIp(router)
	if secondary != "" {
		cmd("ip", "address", "add", secondary, "dev", "eth0")
	}

	//Agent
	c := &con.Context{Cmd: cmd, CmdP: cmdp}
	a := agent{con: c}

	var states *[]st.State
	log.Print("Restarting...")
	state := new(nlan.Model)
	util.GetState(router, state)
	request := nlan.Request{Model: state}
	log.Printf("State for %s: %v", router, state)
	log.Printf("Request: %v", request)
	exit := a.route(env.ADD, &request, configMode)
	log.Printf("Restarted: %d", exit)
}
