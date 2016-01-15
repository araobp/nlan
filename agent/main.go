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
	"github.com/araobp/nlan/agent/rpc"
	"github.com/araobp/nlan/common"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"
	st "github.com/araobp/nlan/state"
	"github.com/araobp/nlan/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

func clear(con *con.Context) {
	rpc.Clear(con)
}

// gRPC Add method
func (a *agent) Add(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	exit := a.route(env.ADD, in, util.CONFIG)
	response := nlan.Response{Exit: exit}
	return &response, nil
}

// gRPC Update method
func (a *agent) Update(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	exit := a.route(env.UPDATE, in, util.CONFIG)
	response := nlan.Response{Exit: exit}
	return &response, nil
}

// gRPC Delete method
func (a *agent) Delete(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	exit := a.route(env.DELETE, in, util.CONFIG)
	response := nlan.Response{Exit: exit}
	return &response, nil
}

// gRPC Hello method
func (a *agent) Hello(ctx context.Context, cp *nlan.Capabilities) (*nlan.Capabilities, error) {
	// TODO: impl
	c := []string{"hello"}
	cs := nlan.Capabilities{Capability: c}
	return &cs, nil
}

// gRPC clear method
func (a *agent) Clear(ctx context.Context, cp *nlan.ClearMode) (*nlan.Response, error) {
	// nlan.ClearMode is ignored.
	clear(a.con)
	response := nlan.Response{Exit: 0, LogMessage: "Cleared"}
	return &response, nil
}

func main() {

	target := os.Getenv("HOSTNAME")
	ope := flag.String("ope", "ADD", "CRUD operation")
	filename := flag.String("state", "", "state file")
	roster := flag.String("roster", "", "roster file")
	modeOption := flag.String("mode", "config", "config mode")
	flag.Parse()

	file := fmt.Sprintf("/var/volume/nlan-agent-%s.log", target)
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fataif("Unable to open log file: %s", file)
	}
	log.SetOutput(f)

	router := util.RegisterHost()
	mode := util.GetMode(router)

	var configMode int
	switch *modeOption {
	case "debug":
		configMode = util.DEBUG
	case "config":
		switch mode {
		case env.RESTART:
			configMode = util.RESTART
		case env.INIT:
			configMode = util.CONFIG
		default:
			configMode = util.CONFIG
		}
	}

	log.Printf("Start mode: %d", mode)
	cmd, cmdp := util.GetCmd(configMode, true)

	//Adds a secondary IP address to eth0
	secondary := util.GetSecondaryIp(router)
	if secondary != "" {
		cmd("ip", "address", "add", secondary, "dev", "eth0")
	}
	c := &con.Context{Cmd: cmd, CmdP: cmdp}
	a := agent{con: c}

	var states *[]st.State
	switch {
	case *filename != "":
		log.Print("### Direct config mode ###")

		switch *roster {
		case "":
			states, _ = common.ReadState(filename, nil)
		default:
			states, _ = common.ReadState(filename, roster)
		}
		for _, v := range *states {
			router := v.Router
			model := v.Model
			log.Printf("target: %s -- router: %s\n", target, router)
			if router == target {
				request := nlan.Request{Model: &model}
				var ope_ int
				switch *ope {
				case "add":
					ope_ = env.ADD
				case "update":
					ope_ = env.UPDATE
				case "delete":
					ope_ = env.DELETE
				}
				a.route(ope_, &request, configMode)
				log.Print(logbuf.String())
			}
		}
	default:
		log.Print("### gRPC server mode ###")

		if mode == env.RESTART {
			log.Print("Restarting...")
			state := new(nlan.Model)
			util.GetState(router, state)
			request := nlan.Request{Model: state}
			log.Printf("State for %s: %v", router, state)
			log.Printf("Request: %v", request)
			exit := a.route(env.ADD, &request, configMode)
			log.Printf("Restarted: %d", exit)
		}
		listen, err := net.Listen("tcp", env.PORT)
		defer listen.Close()
		if err != nil {
			log.Print(err)
		}
		s := grpc.NewServer()
		nlan.RegisterNlanAgentServer(s, &a)
		s.Serve(listen)
	}
}
