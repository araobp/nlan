package ptn

import (
	"bytes"
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
	"github.com/araobp/go-nlan/nlan/util"
	"strconv"
)

func addVpls(sVni string, sVid string, ip string, brInt string, con *context.Context) {
	cmd, cmdp := con.GetCmd()
	_ = con.Logger
	intBr := "int_br" + sVni
	cmdp("ovs-vsctl", "add-port", brInt, intBr, "tag="+sVid, "-- set interface", intBr, "type=internal")
	cmd("ip", "link", "set", "dev", intBr, "up")
	cmd("ip", "addr", "add", "dev", intBr, ip)
}

func updateVpls(vni int, vid int, ip string, brInt string, con *context.Context) {
	//
}

func deleteVpls(vni int, vid int, ip string, brInt string, con *context.Context) {
	//
}

func addFlowEntries(sVid string, sVni string, peers *[]string, brTun string, con *context.Context) {
	cmd, _ := con.GetCmd()
	_ = con.Logger
	_ = "int_br" + sVni
	cmd("ovs-ofctl", "add-flow", brTun, "table=2,priority=1,tun_id="+sVni+",actions=mod_vlan_vid:"+sVid+",resubmit(,10)")
	// Broadcast tree for each vni
	l := util.GetVxlanPorts(peers)
	var buff bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value
		buff.Write([]byte(",output:"))
		buff.Write([]byte(string((v.(int64)))))
	}
	outputPorts := buff.String()
	cmd("ovs-ofctl", "add-flow", brTun, "table=21,priority=1,dl_vlan="+sVid+",actions=strip_vlan,set_tunnel:"+sVni+outputPorts)
}

func AddL2Vpn(l2vpn *nlan.L2Vpn, con *context.Context, brTun string, brInt string) {
	logger := con.Logger
	ip := l2vpn.Ip
	peers := l2vpn.Peers
	vid := l2vpn.Vid
	vni := l2vpn.Vni
	sVid := strconv.FormatUint(uint64(vid), 10)
	sVni := strconv.FormatUint(uint64(vni), 10)
	logger.Printf("Adding vlan: %s", sVid)
	addVpls(sVid, sVni, ip, brInt, con)
	addFlowEntries(sVid, sVni, &peers, brTun, con)
}

func UpdateL2Vpn(l2vpn *nlan.L2Vpn, con *context.Context, brTun string, brInt string) {
}

func DeleteL2Vpn(l2vpn *nlan.L2Vpn, con *context.Context, brTun string, brInt string) {
}
