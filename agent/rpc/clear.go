package rpc

import (
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/util"

	"log"
)

func Clear(con *context.Context) {
	log.Print("Clearing...")
	cmd, _ := con.GetCmd()
	l := util.GetBridgeNames()
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value
		bridge := v.(string)
		cmd("ovs-vsctl", "del-br", bridge)
	}
}
