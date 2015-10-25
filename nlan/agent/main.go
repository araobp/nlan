package main

import (
	"encoding/json"
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

func log_ope(ope int, in *nlan.Request) {
	log.Printf("Server: %v is called\n", ope)
	log.Print(in)
	json_data, _ := json.Marshal(in)
	log.Printf("%s\n", json_data)
}

func route(ope int, in *nlan.Request) {
	model := in.Model
	ptn := model.GetPtn()
	dvr := model.GetDvr()
	if ptn != nil {
		switch ope {
		case env.ADD:
			config_ptn.Add(ptn)
		case env.UPDATE:
			config_ptn.Update(ptn)
		case env.DELETE:
			config_ptn.Delete(ptn)
		}
	}
	if dvr != nil {
		//
	}
}

// Add method
func (a *agent) Add(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	route(env.ADD, in)
	response := nlan.Response{Exit: 0, LogMessage: "Server: Add() is called"}
	return &response, nil
}

// Update method
func (a *agent) Update(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	route(env.UPDATE, in)
	response := nlan.Response{Exit: 0, LogMessage: "Server: Update() is called"}
	return &response, nil
}

// Delete method
func (a *agent) Delete(ctx context.Context, in *nlan.Request) (*nlan.Response, error) {
	route(env.DELETE, in)
	response := nlan.Response{Exit: 0, LogMessage: "Server: Delete() is called"}
	return &response, nil
}

// Hello method
func (a *agent) Hello(ctx context.Context, cp *nlan.Capabilities) (*nlan.Capabilities, error) {
	// TODO: impl
	return nil, nil
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
