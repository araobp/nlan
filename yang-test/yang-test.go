package main

import (
	"fmt"

	br "github.com/araobp/golan/yang-test/model/bridges"
	sn "github.com/araobp/golan/yang-test/model/subnets"
	vx "github.com/araobp/golan/yang-test/model/vxlan"
	"github.com/golang/protobuf/proto"
)

func be(s string) {
	fmt.Printf("Before encoding: %s\n", s)
}

func wf(v []byte) {
	fmt.Printf("Wire format(encoded): %v\n", v)
}

func ad(s string) {
	fmt.Printf("After decoding: %s\n", s)
}

func main() {

	// NLAN bridges module
	enabled := true
	bridges := br.Bridges{OvsBridges: &enabled}
	fmt.Println("--- NLAN bridges module ---")
	be(bridges.String())
	// Encoding
	data, _ := proto.Marshal(&bridges)
	wf(data)
	// Decoding
	newBridges := br.Bridges{}
	_ = proto.Unmarshal(data, &newBridges)
	ad(newBridges.String())
	fmt.Println("")

	// NLAN vxlan modules
	localIp := "10.0.0.1"
	remoteIps := []string{"192.168.56.101", "192.168.56.102"}
	vxlan := vx.Vxlan_{LocalIp: &localIp, RemoteIps: remoteIps}
	fmt.Println("--- NLAN vxlan module ---")
	be(vxlan.String())
	// Encoding
	data, _ = proto.Marshal(&vxlan)
	wf(data)
	// Decoding
	newVxlan := vx.Vxlan_{}
	_ = proto.Unmarshal(data, &newVxlan)
	ad(newVxlan.String())
	fmt.Println("")

	// NLAN subnets module
	addr1 := "172.16.42.1"
	mode1 := "dvr"
	ipDvr1 := sn.IpDvr{Addr: &addr1, Dhcp: nil, Mode: &mode1}
	ipDvr := []*sn.IpDvr{&ipDvr1}
	var vid1 uint32 = 10
	var vni1 uint32 = 100
	subnets := sn.Subnets_{IpDvr: ipDvr, Peers: nil, Ports: nil, Vid: &vid1, Vni: &vni1}
	fmt.Println("--- NLAN subnets module ---")
	be(subnets.String())
	// Encoding
	data, _ = proto.Marshal(&subnets)
	wf(data)
	// Decoding
	newSubnets := sn.Subnets_{}
	_ = proto.Unmarshal(data, &newSubnets)
	ad(newSubnets.String())
	fmt.Println("")
}
