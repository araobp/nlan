package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
)

func Add(in *nlan.Ptn, con *context.Context) {
	networks := in.GetNetworks()
	cmd, cmdp := con.GetCmd()
	logger := con.Logger
	for _, net := range networks {
		logger.Println(net)
		nodes := net.GetNodes()
		brTun := nodes.Ptn
		brInt := nodes.L2Sw
		logger.Println(brTun)
		logger.Println(brInt)
		cmd("date")
		cmdp("date")

	}
}

func Update(in *nlan.Ptn, con *context.Context) {
	logger := con.Logger
	logger.Print(in)
}

func Delete(in *nlan.Ptn, con *context.Context) {
	logger := con.Logger
	logger.Print(in)
}
