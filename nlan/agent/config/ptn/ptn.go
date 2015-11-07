package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
)

func Add(in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	for _, net := range networks {
		con.Logger.Println(net)
		nodes := net.GetNodes()
		AddNodes(nodes, con)
	}
}

func Update(in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	for _, net := range networks {
		con.Logger.Println(net)
		nodes := net.GetNodes()
		UpdateNodes(nodes, con)
	}
}

func Delete(in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	for _, net := range networks {
		con.Logger.Println(net)
		nodes := net.GetNodes()
		DeleteNodes(nodes, con)
	}
}
