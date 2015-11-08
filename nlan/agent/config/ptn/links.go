package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
	"github.com/araobp/go-nlan/nlan/util"
)

func AddLinks(links *nlan.Links, con *context.Context, brTun string, brInt string) {
	_, cmdp := con.GetCmd()
	logger := con.Logger
	localIp := links.LocalIp
	remoteIps := links.RemoteIps
	for _, remoteIp := range remoteIps {
		inf := util.ZfillIp(remoteIp)
		logger.Printf("Adding a VXLAN tunnel: %s", inf)
		cmdp("ovs-vsctl", "add-port", brTun, inf, "--", "set", "interface", inf, "type=vxlan", "options:in_key=flow", "options:local_ip="+localIp, "options:out_key=flow", "options:remote_ip="+remoteIp)
		/*
		   	vxlan_ports = get_vxlan_ports()
		   	for vxlan_port in vxlan_ports:
		       		cmd('ovs-ofctl add-flow', br_tun, 'table=0,priority=1,in_port='+vxlan_port+',actions=resubmit(,2)')
		*/
	}
}

func UpdateLinks(links *nlan.Nodes, con *context.Context, brTun string, brInt string) {
}

func DeleteLinks(links *nlan.Nodes, con *context.Context, brTun string, brInt string) {
}
