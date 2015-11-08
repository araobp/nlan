package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
)

func Add(in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	logger := con.Logger
	for _, net := range networks {
		logger.Println(net)
		nodes := net.GetNodes()
		if nodes == nil {
			logger.Fatal("nodes required")
		}
		links := net.GetLinks()
		if links == nil {
			logger.Fatal("links required")
		}
		l2vpn := net.GetL2Vpn()
		if l2vpn == nil {
			logger.Fatal("l2vpn required")
		}
		brTun, brInt := AddNodes(nodes, con)
		AddLinks(links, con, brTun, brInt)
		for _, vpn := range l2vpn {
			AddL2Vpn(vpn, con)
		}
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
