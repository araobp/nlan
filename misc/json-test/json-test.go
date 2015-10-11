package main

import (
	"encoding/json"
	"fmt"
)

type Subnets struct {
	IpDvr []*IpDvr `protobuf:"bytes,1,rep,name=IpDvr" json:"IpDvr,omitempty"`
	Peers []string `protobuf:"bytes,2,rep,name=Peers" json:"Peers,omitempty"`
	Ports []string `protobuf:"bytes,3,rep,name=Ports" json:"Ports,omitempty"`
	Vid   uint32   `protobuf:"varint,4,opt,name=Vid" json:"Vid,omitempty"`
	Vni   uint32   `protobuf:"varint,5,opt,name=Vni" json:"Vni,omitempty"`
}

type IpDvr struct {
	Addr string `protobuf:"bytes,1,opt,name=Addr" json:"Addr,omitempty"`
	Dhcp string `protobuf:"bytes,2,opt,name=Dhcp" json:"Dhcp,omitempty"`
	Mode string `protobuf:"bytes,3,opt,name=Mode" json:"Mode,omitempty"`
}

func main() {

	byte_ := []byte(`{"vni":1, "vid":10, "peers": ["192.168.56.101", "192.168.56.102"]}`)
	data := Subnets{}

	if err := json.Unmarshal(byte_, &data); err != nil {
		panic(err)
	}
	fmt.Printf("Vid: %v\n", data.Vid)
	fmt.Printf("Vni: %v\n", data.Vni)
	fmt.Printf("Peers: %v\n", data.Peers)
	fmt.Printf("Ports: %v\n", data.Ports)
	fmt.Printf("IpDvr: %v\n", data.IpDvr)
}
