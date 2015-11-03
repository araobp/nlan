package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/env"
	"github.com/araobp/go-nlan/nlan/model/nlan"
)

func ConfNodes(crud int, nodes *nlan.Nodes, con *context.Context) {
	cmd, cmdp := con.GetCmd()
	logger := con.Logger
	brTun := nodes.Ptn
	brInt := nodes.L2Sw
	logger.Println(brTun)
	logger.Println(brInt)
	switch crud {
	case env.ADD:
		logger.Print("nodes Add() called")
		cmd("echo", "cmd")
		cmdp("echo", "cmdp")
	case env.UPDATE:
	case env.DELETE:
	}
}
