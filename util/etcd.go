package util

import (
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/golang/protobuf/proto"
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
func RegisterHost() {

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
}

// Gets a list of all host names and their addresses from etcd
func ListHosts() map[string]interface{} {

	kapi, cont := getKapi()
	key := "/nlan/hosts"
	list, err := kapi.Get(cont, key, &client.GetOptions{Recursive: true})
	if err != nil {
		log.Fatal(err)
	}
	nodes := list.Node.Nodes
	hosts := make(map[string]interface{})
	for _, node := range nodes {
		path := strings.Split(node.Key, "/")
		ipmask := strings.Split(node.Value, "/")
		hostname := path[2]
		ip := ipmask[0]
		hosts[hostname] = ip
	}
	return hosts
}

// Sets NLAN state to etcd
func SetState(hostname string, pb proto.Message) {

	kapi, cont := getKapi()
	key := "/nlan/config/" + hostname
	wire, _ := proto.Marshal(pb)
	value := string(wire)
	_, err := kapi.Set(cont, key, value, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN state from etcd
func GetState(hostname string, pb proto.Message) {

	kapi, cont := getKapi()
	key := "/nlan/config/" + hostname
	r, err := kapi.Get(cont, key, &client.GetOptions{Recursive: false})
	if err != nil {
		log.Fatal(err)
	}
	wire := []byte(r.Node.Value)
	proto.Unmarshal(wire, pb)
}
