package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/env"
	"github.com/araobp/go-nlan/nlan/model/nlan"
	"github.com/araobp/go-nlan/nlan/util"
)

func ConfNodes(crud int, nodes *nlan.Nodes, con *context.Context) {
	cmd, cmdp := con.GetCmd()
	logger := con.Logger
	brTun := nodes.Ptn
	brInt := nodes.L2Sw
	patchTun := "patch-tun_" + brTun
	patchInt := "patch-int_" + brInt
	switch crud {
	case env.ADD:
		logger.Printf("Adding bridges: %s and %s\n", brTun, brInt)
		cmdp("ovs-vsctl", "add-br", brInt)
		cmdp("ovs-vsctl", "add-br", brTun)
		cmd("ovs-ofctl", "del-flows", brTun)
		cmdp("ovs-vsctl", "add-port", brInt, patchInt, "--", "set", "interface", patchInt, "type=patch", "options:peer="+patchTun)
		cmdp("ovs-vsctl", "add-port", brTun, patchTun, "--", "set", "interface", patchTun, "type=patch", "options:peer="+patchInt)
		cmdp("ovs-vsctl", "set-fail-mode", brTun, "secure")
		ofport := util.GetOfport(patchTun)
		logger.Printf("ofport: %d\n", ofport)
	case env.UPDATE:
	case env.DELETE:
	}
}
