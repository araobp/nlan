package main

import (
	"fmt"
)

type Vxlan struct {
	addresses []Addresses
}

type Addresses struct {
	local_ip   string
	remote_ips []string
}

func main() {
	addrs := []Addresses{
		Addresses{
			local_ip: "10.0.0.1",
			remote_ips: []string{
				"192.168.56.101",
				"192.168.56.102"}},
		Addresses{
			local_ip: "10.0.0.2",
			remote_ips: []string{
				"192.168.56.103",
				"192.168.56.104"}},
	}
	vxlan := Vxlan{addresses: addrs}
	fmt.Println(vxlan)
}
