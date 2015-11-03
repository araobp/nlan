package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/env"
	"github.com/araobp/go-nlan/nlan/model/nlan"
)

func Add(in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	for _, net := range networks {
		con.Logger.Println(net)
		nodes := net.GetNodes()
		ConfNodes(env.ADD, nodes, con)
	}
}

func Update(in *nlan.Ptn, con *context.Context) {
	logger := con.Logger
	logger.Print(in)
}

func Delete(in *nlan.Ptn, con *context.Context) {
	logger := con.Logger
	logger.Print(in)
}
