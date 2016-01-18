package util

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/araobp/nlan/model/nlan"
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
func listHosts(path string) map[string]interface{} {
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

// Lists up all hosts on nlan.hosts
func ListHosts() map[string]interface{} {
	return listHosts("nlan.hosts")
}

// Lists up all hosts on nlan.ip
func ListIps() map[string]interface{} {
	return listHosts("nlan.ip")
}

// Sets NLAN state to tega
func SetModel(hostname string, model *nlan.Model) {
	path := fmt.Sprintf("nlan.state.%s", hostname)
	err := ope.Put(path, model)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN state from tega
func GetModel(hostname string, model *nlan.Model) {
	path := fmt.Sprintf("nlan.state.%s", hostname)
	err := ope.Get(path, model)
	if err != nil {
		log.Fatal(err)
	}
}

// Resets NLAN state on tega
func ResetState() {
	err := ope.Delete("nlan.state")
	if err != nil {
		log.Print(err)
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
