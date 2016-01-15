package util

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/araobp/nlan/env"
	"github.com/araobp/tega/driver"
)

var ope *driver.Operation
var hostname string

type Self struct {
}

func (r *Self) OnNotify(v *[]driver.Notify) {
}

func init() {
	var err error
	hostname = os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}
	tega := os.Getenv("TEGA_ADDRESS")
	self := &Self{}
	ope, err = driver.NewOperation(hostname, tega, 0, self)
	if err != nil {
		log.Fatal(err)
	}
}

// Registers a host IP address with tega
func RegisterHost() string {
	path := "nlan.hosts." + hostname
	interfaces, _ := net.Interfaces()
	var addrs []net.Addr
	for _, inter := range interfaces {
		if inter.Name == "eth0" {
			addrs, _ = inter.Addrs()
		}
	}
	value := addrs[0].String()
	err := ope.Put(path, value)
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}

// Gets a list of all host names and their addresses from tega
func ListHosts(secondary bool) map[string]interface{} {
	var path string
	switch secondary {
	case false:
		path = "nlan.hosts"
	case true:
		path = "nlan.ip"
	}
	var nodes map[string]interface{}
	err := ope.Get(path, &nodes)
	if err != nil {
		log.Fatal(err)
	}
	hosts := make(map[string]interface{})
	for host, ipmask := range nodes {
		log.Print(host)
		log.Print(ipmask)
		hosts[host] = strings.Split(ipmask.(string), "/")[0]
	}
	return hosts
}

// Sets NLAN state to tega
func SetState(hostname string, state interface{}) {
	path := fmt.Sprintf("nlan.state.%s.json", hostname)
	err := ope.Put(path, &state)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN state from tega
func GetState(hostname string, state interface{}) {
	path := fmt.Sprintf("nlan.state.%s.json", hostname)
	err := ope.Get(path, &state)
	if err != nil {
		log.Fatal(err)
	}
}

// Sets NLAN mode to tega
func SetMode(hostname string, mode int) {
	path := fmt.Sprintf("nlan.state.%s.mode", hostname)
	err := ope.Put(path, &mode)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN mode from tega
func GetMode(hostname string) int {
	var i int
	var mode int
	path := fmt.Sprintf("nlan.state.%s.mode", hostname)
	err := ope.Get(path, &mode)
	if err != nil {
		i = env.INIT
	} else {
		i = mode
	}
	return i
}

// Resets NLAN state on tega
func ResetState() {
	path := "nlan.state"
	var routers map[string]interface{}
	err := ope.Get(path, &routers)
	if err == nil {
		for router, _ := range routers {
			ope.Delete(fmt.Sprintf("nlan.state.%s", router))
		}
	}
}

// Gets a secondary IP address from tega
func GetSecondaryIp(hostname string) string {
	path := "nlan.ip." + hostname
	var secondary string
	err := ope.Get(path, &secondary)
	if err != nil {
		log.Fatal(err)
	}
	return secondary
}
