package util

import (
	netstat "github.com/drael/GOnetstat"
)

const (
	TCP = "TCP"
	UDP = "UDP"
)

func Netstat(t string) *[]netstat.Process {
	var data []netstat.Process
	switch t {
	case TCP:
		data = netstat.Tcp()
	case UDP:
		data = netstat.Udp()
	}
	return &data
}
