package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
	_ "github.com/araobp/go-nlan/nlan/util"
)

func AddL2Vpn(nodes *nlan.L2Vpn, con *context.Context) {
	_, _ = con.GetCmd()
	_ = con.Logger
}

func UpdateL2Vpn(nodes *nlan.L2Vpn, con *context.Context) {
}

func DeleteL2Vpn(nodes *nlan.L2Vpn, con *context.Context) {
}
