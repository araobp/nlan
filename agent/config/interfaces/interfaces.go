package interfaces

import (
	"github.com/araobp/nlan/agent/context"
	"github.com/araobp/nlan/model/nlan"

	"log"
)

func Crud(crud int, in map[string]*nlan.Interface, con *context.Context) {
	// TODO: CRUD operatoins
	cmd, _ := con.GetCmd()
	for dev, params := range in { // map key is not used.
		log.Println(dev)
		cmd("ip", "tunnel", "add", dev, "mode", params.Mode, "local", params.Local, "remote", params.Remote)
		cmd("ip", "addr", "add", params.Address, "dev", dev)
		cmd("ip", "link", "set", "dev", dev, "up")
	}
}
