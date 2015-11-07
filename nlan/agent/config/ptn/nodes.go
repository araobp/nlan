package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
	"github.com/araobp/go-nlan/nlan/util"
)

func AddNodes(nodes *nlan.Nodes, con *context.Context) {
	cmd, cmdp := con.GetCmd()
	logger := con.Logger
	brTun := nodes.Ptn
	brInt := nodes.L2Sw
	patchTun := "patch-tun_" + brTun
	patchInt := "patch-int_" + brInt
	logger.Printf("Adding bridges: %s and %s\n", brTun, brInt)
	// Adds br-int and br-tun and connects them to each other
	cmdp("ovs-vsctl", "add-br", brInt)
	cmdp("ovs-vsctl", "add-br", brTun)
	cmd("ovs-ofctl", "del-flows", brTun)
	cmdp("ovs-vsctl", "add-port", brInt, patchInt, "--", "set", "interface", patchInt, "type=patch", "options:peer="+patchTun)
	cmdp("ovs-vsctl", "add-port", brTun, patchTun, "--", "set", "interface", patchTun, "type=patch", "options:peer="+patchInt)
	cmdp("ovs-vsctl", "set-fail-mode", brTun, "secure")
	// Obtains ofport for 'patch-tun' port
	patchTunNum := string(util.GetOfport(patchTun))
	logger.Printf("patchTunNum(ofport): %s\n", patchTunNum)
	// Adds flow entries onto br-tun
	cmd("ovs-ofctl", "add-flow", brTun, "table=0,priority=1,in_port="+patchTunNum+",actions=resubmit(,1)")
	cmd("ovs-ofctl", "add-flow", brTun, "table=0,priority=0,actions=drop")
	cmd("ovs-ofctl", "add-flow", brTun, "table=1,priority=0,dl_dst=01:00:00:00:00:00/01:00:00:00:00:00,actions=resubmit(,19)")
	cmd("ovs-ofctl", "add-flow", brTun, "table=1,priority=0,dl_dst=00:00:00:00:00:00/01:00:00:00:00:00,actions=resubmit(,20)")
	cmd("ovs-ofctl", "add-flow", brTun, "table=2,priority=0,actions=drop")
	cmd("ovs-ofctl", "add-flow", brTun, "table=3,priority=0,actions=drop")
	cmd("ovs-ofctl", "add-flow", brTun, "table=10,priority=1,actions=learn(table=20,hard_timeout=300,priority=1,NXM_OF_VLAN_TCI[0..11],NXM_OF_ETH_DST[]=NXM_OF_ETH_SRC[],load:0->NXM_OF_VLAN_TCI[],load:NXM_NX_TUN_ID[]->NXM_NX_TUN_ID[],output:NXM_OF_IN_PORT[]),output:"+patchTunNum)
	cmd("ovs-ofctl", "add-flow", brTun, "table=19,priority=0,actions=resubmit(,21)")
	cmd("ovs-ofctl", "add-flow", brTun, "table=20,priority=0,actions=resubmit(,21)")
	cmd("ovs-ofctl", "add-flow", brTun, "table=21,priority=0,actions=drop")
}

func UpdateNodes(nodes *nlan.Nodes, con *context.Context) {
}

func DeleteNodes(nodes *nlan.Nodes, con *context.Context) {
}
