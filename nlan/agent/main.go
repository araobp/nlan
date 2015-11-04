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

func route(ope int, in *nlan.Request) string {
	model := in.Model
	log.Print(model)
	ptn := model.GetPtn()
	dvr := model.GetDvr()
	var logbuf bytes.Buffer
	logger := log.New(&logbuf, "", log.LstdFlags)
	cmd, cmdp := util.GetCmd(logger, util.DEBUG)
	con := &con.Context{Cmd: cmd, CmdP: cmdp, Logger: logger}
	if ptn != nil {
		switch ope {
		case env.ADD:
			logger.Print("--ADD")
			config_ptn.Add(ptn, con)
		case env.UPDATE:
			logger.Print("--UPDATE")
			config_ptn.Update(ptn, con)
		case env.DELETE:
			logger.Print("--DELETE")
			config_ptn.Delete(ptn, con)
		}
	}
	if dvr != nil {
		//
	}
	return logbuf.String()
}

// Add method
func (a *agent) Add(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	logMessage := route(env.ADD, in)
	response := nlan.Response{Exit: 0, LogMessage: logMessage}
	return &response, nil
}

// Update method
func (a *agent) Update(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	logMessage := route(env.UPDATE, in)
	response := nlan.Response{Exit: 0, LogMessage: logMessage}
	return &response, nil
}

// Delete method
func (a *agent) Delete(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	logMessage := route(env.DELETE, in)
	response := nlan.Response{Exit: 0, LogMessage: logMessage}
	return &response, nil
}

// Hello method
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
	flag.Parse()
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
				logMessage := route(ope_, &request)
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
