package router

import (
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"
	"github.com/araobp/nlan/util"
)

func Crud(crud int, in *nlan.Router, con *context.Context) {

	loopback := in.Loopback
	ospf := in.GetOspf()
	logger := con.Logger
	logger.Print("Router called...")
	var crudRouter func(string, []*nlan.Ospf, *context.Context)

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
	crudRouter(loopback, ospf, con)
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

func addRouter(loopback string, ospf []*nlan.Ospf, con *context.Context) {

	cmd, cmdp := con.GetCmd()
	logger := con.Logger
	cmd("ip", "addr", "add", "dev", "lo", loopback)
	if ospf != nil {
		var script [][]string
		script = append(script, []string{"router", "ospf"})
		script = append(script, []string{"redistributed", "connected"})
		for _, o := range ospf {
			area := o.Area
			networks := o.Networks
			logger.Print("Networks: %v", networks)
			routerOspfNetworks(&script, area, networks)
		}
		script = append(script, []string{"exit"})
		batch := util.VtyshBatch(script)
		cmdp("vtysh", batch...)
	}
}

func updateRouter(loopback string, ospf []*nlan.Ospf, con *context.Context) {
	//
}

func deleteRouter(loopback string, ospf []*nlan.Ospf, con *context.Context) {
	//
}
