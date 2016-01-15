package ptn

import (
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"

	"log"
)

func Crud(crud int, in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	var crudNodes func(*nlan.Nodes, *context.Context) (string, string)
	var crudLinks func(*nlan.Links, *context.Context, string, string)
	var crudL2Vpn func(*nlan.L2Vpn, *context.Context, string, string) (string, string, string)
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
	default:
		log.Fatal("CRUD unidentified")
	}
	for _, net := range networks {
		log.Println(net)
		nodes := net.GetNodes()
		if nodes == nil {
			log.Fatal("nodes required")
		}
		links := net.GetLinks()
		if links == nil {
			log.Fatal("links required")
		}
		l2vpn := net.GetL2Vpn()
		if l2vpn == nil {
			log.Fatal("l2vpn required")
		}

		log.Printf("L2Sw: %s, Ptn: %s", nodes.L2Sw, nodes.Ptn)
		brTun, brInt := crudNodes(nodes, con)
		log.Printf("crudNodes() completed")

		crudLinks(links, con, brTun, brInt)
		for _, vpn := range l2vpn {
			ip, sVid, sVni := crudL2Vpn(vpn, con, brTun, brInt)
			log.Printf("crudL2Vpn() completed: %s, %s, %s", ip, sVid, sVni)
		}
	}
}
