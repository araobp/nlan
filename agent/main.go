package main

import (
	"fmt"
	"log"
	"os"
	"time"

	config_interfaces "github.com/araobp/nlan/agent/config/interfaces"
	config_ptn "github.com/araobp/nlan/agent/config/ptn"
	config_router "github.com/araobp/nlan/agent/config/router"
	config_vhosts "github.com/araobp/nlan/agent/config/vhosts"
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"
	"github.com/araobp/nlan/util"
)

type agent struct {
	con *context.Context
}

func (a *agent) route(crud int, model *nlan.Model) (exit uint32) {
	ptn := model.GetPtn()
	dvr := model.GetDvr()
	vhosts := model.GetVhosts()
	router := model.GetRouter()
	interfaces := model.GetInterfaces()

	exit = 0
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			exit = 1
		}
	}()

	if ptn != nil {
		log.Print("Routing to ptn module...")
		config_ptn.Crud(crud, ptn, a.con)
	}
	if interfaces != nil {
		log.Print("Routing to interfaces module...")
		config_interfaces.Crud(crud, interfaces, a.con)
	}
	if dvr != nil {
		//
	}
	if vhosts != nil {
		log.Print("Routing to vhosts module...")
		config_vhosts.Crud(crud, vhosts, a.con)
	}
	if router != nil {
		log.Print("Routing to router module...")
		config_router.Crud(crud, router, a.con)
	}
	return exit
}

func main() {

	// Command options
	target := os.Getenv("HOSTNAME")

	// Log file setup
	file := fmt.Sprintf("/var/volume/nlan-agent-%s.log", target)
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Unable to open log file: %s", file)
	}
	log.SetOutput(f)

	// Gets cmd and cmdp
	cmd, cmdp := util.GetCmd(true)

	//Adds a secondary IP address to eth0
	router := util.GetHostname()
	secondary := util.GetSecondaryIp(router)
	if secondary != "" {
		cmd("ip", "address", "add", secondary, "dev", "eth0")
	}

	//Agent
	c := &context.Context{Cmd: cmd, CmdP: cmdp}
	a := agent{con: c}

	log.Print("Starting...")
	model := new(nlan.Model)
	util.GetModel(router, model)
	log.Printf("State for %s: %v", router, model)
	exit := a.route(env.ADD, model)
	log.Printf("Started: %d", exit)

	//Infinite loop
	for {
		time.Sleep(1 * time.Second)
	}
}
