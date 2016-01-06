package router

import (
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/env"
	"github.com/araobp/nlan/model/nlan"
	"github.com/araobp/nlan/util"
	api "github.com/osrg/gobgp/api"
	"github.com/osrg/gobgp/packet"
	gobgp "github.com/osrg/gobgp/server"

	"strconv"
	"strings"
)

func Crud(crud int, in *nlan.Router, con *context.Context) {

	loopback := in.Loopback
	embedded := in.EmbeddedBgp // quagga-bgpd(false) or gobgp(true)
	ospf := in.GetOspf()
	bgp := in.GetBgp()
	cmd, _ := con.GetCmd()
	logger := con.Logger
	logger.Print("Router called...")

	var s *gobgp.BgpServer
	var g *gobgp.Server
	if embedded {
		cmd("/etc/init.d/quagga", "stop")
		s = gobgp.NewBgpServer()
		go s.Serve()
		g = gobgp.NewGrpcServer(50051, s.GrpcReqCh)
		go g.Serve()
	}

	var crudRouter func(string, bool, *gobgp.BgpServer, []*nlan.Ospf, []*nlan.Bgp, *context.Context)

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
	crudRouter(loopback, embedded, s, ospf, bgp, con)
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

func gobgpReqModNeighbor(s *gobgp.BgpServer, neighs []*nlan.Neighbors, con *context.Context) {
	logger := con.Logger
	for _, n := range neighs {
		peer := n.Peer
		as := n.RemoteAs
		client := n.RouteReflectorClient
		p := api.Peer{}
		p.Conf = &api.PeerConf{
			NeighborAddress: peer,
			PeerAs:          as,
		}
		if client == true {
			p.RouteReflector = &api.RouteReflector{
				RouteReflectorClient: true,
			}
		}
		req := gobgp.NewGrpcRequest(gobgp.REQ_MOD_NEIGHBOR, "", bgp.RouteFamily(0), &api.ModNeighborArguments{
			Operation: api.Operation_ADD,
			Peer:      &p,
		})
		s.GrpcReqCh <- req
		res := <-req.ResponseCh
		if err := res.Err(); err != nil {
			logger.Print(err)
		}
	}
}

func gobgpReqModGlobalConfig(s *gobgp.BgpServer, routerId string, as int64, con *context.Context) {
	logger := con.Logger
	req := gobgp.NewGrpcRequest(gobgp.REQ_MOD_GLOBAL_CONFIG, "", bgp.RouteFamily(0), &api.ModGlobalConfigArguments{
		Operation: api.Operation_ADD,
		Global: &api.Global{
			As:         uint32(as),
			RouterId:   routerId,
			ListenPort: -1, // gobgp won't listen on tcp:179
		},
	})
	s.GrpcReqCh <- req
	res := <-req.ResponseCh
	if err := res.Err(); err != nil {
		logger.Print(err)
	}
}

func addRouter(loopback string, embedded bool, s *gobgp.BgpServer, ospf []*nlan.Ospf, bgp []*nlan.Bgp, con *context.Context) {

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
			if embedded {
				routerId := strings.Split(loopback, "/")[0]
				gobgpReqModGlobalConfig(s, routerId, int64(b.As), con)
			} else {
				script = append(script, []string{"router", "bgp", strconv.FormatUint(uint64(b.As), 10)})
				script = append(script, []string{"redistribute", "connected"})
			}
			neigh := b.GetNeighbors()
			if neigh != nil {
				if embedded {
					gobgpReqModNeighbor(s, neigh, con)
				} else {
					routerBgpNeighbors(&script, neigh)
				}
			}
		}
		script = append(script, []string{"exit"})
	}
	if len(script) > 0 && !embedded {
		batch := util.VtyshBatch(script)
		cmdp("vtysh", batch...)
	}
}

func updateRouter(loopback string, embedded bool, s *gobgp.BgpServer, ospf []*nlan.Ospf, bgp []*nlan.Bgp, con *context.Context) {
	//
}

func deleteRouter(loopback string, embedded bool, s *gobgp.BgpServer, ospf []*nlan.Ospf, bgp []*nlan.Bgp, con *context.Context) {
	//
}
