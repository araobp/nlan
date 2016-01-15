package util

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/araobp/nlan/env"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// etcd client
func getKapi() (client.KeysAPI, context.Context) {

	etcdAddress := os.Getenv("ETCD_ADDRESS")
	if etcdAddress == "" {
		log.Fatalf("ETCD_ADDRESS unset")
	}

	config := client.Config{
		Endpoints:               []string{etcdAddress},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	cont := context.Background()
	c, err := client.New(config)
	if err != nil {
		log.Fatal(err)
	}
	return client.NewKeysAPI(c), cont
}

// Registers a host IP address with etcd
func RegisterHost() string {

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		log.Fatalf("HOSTNAME unset")
	}
	kapi, cont := getKapi()

	key := "/nlan/hosts/" + hostname
	interfaces, _ := net.Interfaces()
	var addrs []net.Addr
	for _, inter := range interfaces {
		if inter.Name == "eth0" {
			addrs, _ = inter.Addrs()
		}
	}
	value := addrs[0].String()
	_, err := kapi.Set(cont, key, value, nil)
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}

// Gets a list of all host names and their addresses from etcd
func ListHosts(secondary bool) map[string]interface{} {

	kapi, cont := getKapi()
	var key string
	switch secondary {
	case false:
		key = "/nlan/hosts"

	case true:
		key = "/nlan/ip"
	}
	list, err := kapi.Get(cont, key, &client.GetOptions{Recursive: true})
	hosts := make(map[string]interface{})
	if err == nil {
		nodes := list.Node.Nodes
		for _, node := range nodes {
			path := strings.Split(node.Key, "/")
			ipmask := strings.Split(node.Value, "/")
			hostname := path[3]
			ip := ipmask[0]
			hosts[hostname] = ip
		}
	}
	return hosts
}

// Sets NLAN state to etcd
func SetState(hostname string, state interface{}) {

	kapi, cont := getKapi()
	key := "/nlan/state/" + hostname + "/json"
	wire, _ := json.Marshal(state)
	value := string(wire)
	_, err := kapi.Set(cont, key, value, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN state from etcd
func GetState(hostname string, state interface{}) {

	kapi, cont := getKapi()
	key := "/nlan/state/" + hostname + "/json"
	r, err := kapi.Get(cont, key, &client.GetOptions{Recursive: true})
	if err == nil {
		wire := []byte(r.Node.Value)
		err = json.Unmarshal(wire, state)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Sets NLAN mode to etcd
func SetMode(hostname string, mode int) {

	kapi, cont := getKapi()
	key := "/nlan/state/" + hostname + "/mode"
	value := strconv.Itoa(mode)
	_, err := kapi.Set(cont, key, value, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN mode from etcd
func GetMode(hostname string) int {

	var i int
	kapi, cont := getKapi()
	key := "/nlan/state/" + hostname + "/mode"
	r, err := kapi.Get(cont, key, &client.GetOptions{Recursive: false})
	if err != nil {
		i = env.INIT
	} else {
		i, _ = strconv.Atoi(r.Node.Value)
	}
	return i
}

// Resets NLAN state on etcd
func ResetState() {

	kapi, cont := getKapi()
	key := "/nlan/state"
	list, err := kapi.Get(cont, key, &client.GetOptions{Recursive: false})
	if err == nil {
		nodes := list.Node.Nodes
		for _, node := range nodes {
			kapi.Delete(cont, node.Key, &client.DeleteOptions{Recursive: true})
		}
	}
}

// Gets a secondary IP address from etcd
func GetSecondaryIp(hostname string) string {

	kapi, cont := getKapi()
	key := "/nlan/ip/" + hostname
	var secondary string
	r, err := kapi.Get(cont, key, &client.GetOptions{Recursive: false})
	if err == nil {
		secondary = r.Node.Value
	}
	return secondary
}
