package main

import (
	"bytes"
	"flag"
	"log"
	"net"
	"os"

	config_ptn "github.com/araobp/go-nlan/nlan/agent/config/ptn"
	con "github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/common"
	env "github.com/araobp/go-nlan/nlan/env"
	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
	st "github.com/araobp/go-nlan/nlan/state"
	"github.com/araobp/go-nlan/nlan/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type agent struct{}

func route(crud int, in *nlan.Request, configMode int) string {
	model := in.Model
	log.Print(model)
	ptn := model.GetPtn()
	dvr := model.GetDvr()
	var logbuf bytes.Buffer
	logger := log.New(&logbuf, "", log.LstdFlags)
	cmd, cmdp := util.GetCmd(logger, configMode)
	con := &con.Context{Cmd: cmd, CmdP: cmdp, Logger: logger}
	if ptn != nil {
		config_ptn.Crud(crud, ptn, con)
	}
	if dvr != nil {
		//
	}
	return logbuf.String()
}

// gRPC Add method
func (a *agent) Add(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	logMessage := route(env.ADD, in, util.CONFIG)
	response := nlan.Response{Exit: 0, LogMessage: logMessage}
	return &response, nil
}

// gRPC Update method
func (a *agent) Update(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	logMessage := route(env.UPDATE, in, util.CONFIG)
	response := nlan.Response{Exit: 0, LogMessage: logMessage}
	return &response, nil
}

// gRPC Delete method
func (a *agent) Delete(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	logMessage := route(env.DELETE, in, util.CONFIG)
	response := nlan.Response{Exit: 0, LogMessage: logMessage}
	return &response, nil
}

// gRPC Hello method
func (a *agent) Hello(ctx context.Context, cp *nlan.Capabilities) (*nlan.Capabilities, error) {
	// TODO: impl
	c := []string{"hello"}
	cs := nlan.Capabilities{Capability: c}
	return &cs, nil
}

func main() {
	target := os.Getenv("HOSTNAME")
	ope := flag.String("ope", "ADD", "CRUD operation")
	filename := flag.String("state", "", "state file")
	roster := flag.String("roster", "", "roster file")
	mode := flag.String("mode", "debug", "config mode")
	flag.Parse()

	var configMode int
	switch *mode {
	case "debug":
		configMode = util.DEBUG
	case "config":
		configMode = util.CONFIG
	case "restart":
		configMode = util.RESTART
	}

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
				logMessage := route(ope_, &request, configMode)
				log.Print(logMessage)
			}
		}
	default:
		log.Print("### gRPC server mode ###")
		util.RegisterHost()
		listen, err := net.Listen("tcp", env.PORT)
		if err != nil {
			log.Print(err)
		}
		s := grpc.NewServer()
		nlan.RegisterNlanAgentServer(s, &agent{})
		s.Serve(listen)
	}
}
