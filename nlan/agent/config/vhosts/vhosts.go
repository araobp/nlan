package vhosts

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/env"
	"github.com/araobp/go-nlan/nlan/model/nlan"

	"strconv"
	"strings"
)

type ipAddress struct {
	network string
	abcd    string
	abc     string
	d       string
	mask    string
}

func getIpAddress(network string) ipAddress {
	a := strings.Split(network, "/")
	abcd := a[0]
	mask := a[1]
	as := strings.Split(abcd, ".")
	abc := strings.Join(as[0:3], ".")
	d := as[3]
	addr := ipAddress{
		network: network,
		abcd:    abcd,
		abc:     abc,
		d:       d,
		mask:    mask,
	}
	return addr
}

func Crud(crud int, in *nlan.Vhosts, con *context.Context) {

	vhostProps := in.GetVhostProps()
	logger := con.Logger
	logger.Print("Vhosts called...")
	var crudVhosts func(ipAddress, uint32, *context.Context)

	switch crud {
	case env.ADD:
		crudVhosts = addVhosts
	case env.UPDATE:
		crudVhosts = updateVhosts
	case env.DELETE:
		crudVhosts = deleteVhosts
	default:
		logger.Fatal("CRUD unidentified")
	}

	for _, props := range vhostProps {
		logger.Print(props)
		network := props.Network
		address := getIpAddress(network)
		vhosts := props.Vhosts

		logger.Printf("Address: %v, Vhosts: %d", address, vhosts)
		crudVhosts(address, vhosts, con)
		logger.Print("crudVhosts() completed")
	}
}

func addVhosts(addr ipAddress, vhosts uint32, con *context.Context) {

	cmd, _ := con.GetCmd()
	_ = con.Logger
	br := "br_" + addr.abcd
	br_ip := addr.abcd
	cmd("brctl", "addbr", br)
	cmd("ip", "addr", "add", "dev", br, addr.network)
	cmd("ip", "link", "set", "dev", br, "up")
	for i := 1; i <= int(vhosts); i++ {
		d, _ := strconv.Atoi(addr.d)
		id := strconv.Itoa(d + i)
		ip := addr.abc + "." + id
		ns := ip
		cmd("ip", "netns", "add", ns)
		cmd("ip", "link", "add", ns, "type", "veth", "peer", "name", "temp")
		cmd("ip", "link", "set", "dev", "temp", "netns", ns)
		cmd("ip", "netns", "exec", ns, "ip", "link", "set", "dev", "temp", "name", "eth0")
		cmd("brctl", "addif", br, ns)
		cmd("ip", "link", "set", "dev", ns, "promisc", "on")
		cmd("ip", "link", "set", "dev", ns, "up")
		cmd("ip", "netns", "exec", ns, "ip", "link", "set", "dev", "eth0", "up")
		cmd("ip", "netns", "exec", ns, "ip", "addr", "add", "dev", "eth0", ip+"/"+addr.mask)
		cmd("ip", "netns", "exec", ns, "ip", "route", "add", "default", "via", br_ip, "dev", "eth0")
	}
}

func updateVhosts(addr ipAddress, vhosts uint32, con *context.Context) {
	//
}

func deleteVhosts(addr ipAddress, vhosts uint32, con *context.Context) {
	//
}
