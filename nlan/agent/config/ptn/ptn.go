package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/env"
	"github.com/araobp/go-nlan/nlan/model/nlan"
)

func Crud(crud int, in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	logger := con.Logger
	var crudNodes func(*nlan.Nodes, *context.Context) (string, string)
	var crudLinks func(*nlan.Links, *context.Context, string, string)
	var crudL2Vpn func(*nlan.L2Vpn, *context.Context, string, string)
	switch crud {
	case env.ADD:
		crudNodes = AddNodes
		crudLinks = AddLinks
		crudL2Vpn = AddL2Vpn
	case env.UPDATE:
		crudNodes = UpdateNodes
		crudLinks = UpdateLinks
		crudL2Vpn = UpdateL2Vpn
	case env.DELETE:
		crudNodes = DeleteNodes
		crudLinks = DeleteLinks
		crudL2Vpn = DeleteL2Vpn
	}
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
		brTun, brInt := crudNodes(nodes, con)
		crudLinks(links, con, brTun, brInt)
		for _, vpn := range l2vpn {
			crudL2Vpn(vpn, con, brTun, brInt)
		}
	}
}
