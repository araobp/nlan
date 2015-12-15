package router

import (
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"
	"github.com/araobp/nlan/util"

	"strconv"
)

func Crud(crud int, in *nlan.Router, con *context.Context) {

	loopback := in.Loopback
	ospf := in.GetOspf()
	bgp := in.GetBgp()
	logger := con.Logger
	logger.Print("Router called...")
	var crudRouter func(string, []*nlan.Ospf, []*nlan.Bgp, *context.Context)

	switch crud {
	case env.ADD:
		crudRouter = addRouter
	case env.UPDATE:
		crudRouter = updateRouter
	case env.DELETE:
		crudRouter = deleteRouter
	default:
		logger.Fatal("CRUD unidentified")
	}

	logger.Printf("Loopback: %s", loopback)
	crudRouter(loopback, ospf, bgp, con)
	logger.Print("crudRouter() completed")
}

func routerOspfNetworks(s *[][]string, area string, networks []string) {
	for _, network := range networks {
		n := []string{}
		n = append(n, "network")
		n = append(n, network)
		n = append(n, "area")
		n = append(n, area)
		*s = append(*s, n)
	}
}

func routerBgpNeighbors(s *[][]string, neighs []*nlan.Neighbors) {
	for _, n := range neighs {
		peer := n.Peer
		as := n.RemoteAs
		client := n.RouteReflectorClient
		nextHopSelf := n.NextHopSelf
		n := []string{}
		n = append(n, "neighbor")
		n = append(n, peer)
		n = append(n, "remote-as")
		n = append(n, strconv.FormatUint(uint64(as), 10))
		*s = append(*s, n)
		if client == true {
			c := []string{}
			c = append(c, "neighbor")
			c = append(c, peer)
			c = append(c, "route-reflector-client")
			*s = append(*s, c)
		}
		if nextHopSelf == true {
			n := []string{}
			n = append(n, "neighbor")
			n = append(n, peer)
			n = append(n, "next-hop-self")
			*s = append(*s, n)
		}
	}
}

func addRouter(loopback string, ospf []*nlan.Ospf, bgp []*nlan.Bgp, con *context.Context) {

	cmd, cmdp := con.GetCmd()
	logger := con.Logger

	// Loopback address
	cmd("ip", "addr", "add", "dev", "lo", loopback)

	// Allow receiving packets from non-best-path interfaces
	cmd("sysctl", "-w", "net.ipv4.conf.all.rp_filter=2")
	// Allow routing packets with local addresses
	cmd("sysctl", "-w", "net.ipv4.conf.all.accept_local=1")

	var script [][]string
	if ospf != nil {
		script = append(script, []string{"router", "ospf"})
		script = append(script, []string{"redistribute", "connected"})
		for _, o := range ospf {
			area := o.Area
			networks := o.Networks
			logger.Print("OSPF Networks: %v", networks)
			routerOspfNetworks(&script, area, networks)
		}
		script = append(script, []string{"exit"})
	}
	if bgp != nil {
		for _, b := range bgp {
			as := b.As
			script = append(script, []string{"router", "bgp", strconv.FormatUint(uint64(as), 10)})
			script = append(script, []string{"redistribute", "connected"})
			neigh := b.GetNeighbors()
			if neigh != nil {
				routerBgpNeighbors(&script, neigh)
			}
		}
		script = append(script, []string{"exit"})
	}
	if len(script) > 0 {
		batch := util.VtyshBatch(script)
		cmdp("vtysh", batch...)
	}
}

func updateRouter(loopback string, ospf []*nlan.Ospf, bgp []*nlan.Bgp, con *context.Context) {
	//
}

func deleteRouter(loopback string, ospf []*nlan.Ospf, bgp []*nlan.Bgp, con *context.Context) {
	//
}
