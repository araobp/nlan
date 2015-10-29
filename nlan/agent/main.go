package main

import (
	"bytes"
	"log"
	"net"

	config_ptn "github.com/araobp/go-nlan/nlan/agent/config/ptn"
	env "github.com/araobp/go-nlan/nlan/env"
	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
	"github.com/araobp/go-nlan/nlan/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type agent struct{}

func route(ope int, in *nlan.Request) string {
	model := in.Model
	ptn := model.GetPtn()
	dvr := model.GetDvr()
	var logbuf bytes.Buffer
	logger := log.New(&logbuf, "", log.LstdFlags)
	if ptn != nil {
		switch ope {
		case env.ADD:
			logger.Print("--ADD")
			config_ptn.Add(ptn, logger)
		case env.UPDATE:
			logger.Print("--UPDATE")
			config_ptn.Update(ptn, logger)
		case env.DELETE:
			logger.Print("--DELETE")
			config_ptn.Delete(ptn, logger)
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
	util.RegisterHost()
	listen, err := net.Listen("tcp", env.PORT)
	if err != nil {
		log.Print(err)
	}
	s := grpc.NewServer()
	nlan.RegisterNlanAgentServer(s, &agent{})
	s.Serve(listen)
}
