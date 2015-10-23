package util

import (
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func RegisterHost() {

	etcdAddress := os.Getenv("ETCD_ADDRESS")
	hostname := os.Getenv("HOSTNAME")
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
	kapi := client.NewKeysAPI(c)

	key := "/hosts/" + hostname
	interfaces, _ := net.Interfaces()
	var addrs []net.Addr
	for _, inter := range interfaces {
		if inter.Name == "eth0" {
			addrs, _ = inter.Addrs()
		}
	}
	value := addrs[0].String()
	_, err = kapi.Set(cont, key, value, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ListHosts() map[string]string {

	etcdAddress := os.Getenv("ETCD_ADDRESS")
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
	kapi := client.NewKeysAPI(c)

	key := "/hosts"
	list, err := kapi.Get(cont, key, &client.GetOptions{Recursive: true})
	if err != nil {
		log.Fatal(err)
	}
	nodes := list.Node.Nodes
	hosts := make(map[string]string)
	for _, node := range nodes {
		ip := strings.Split(node.Value, "/")
		host := ip[0]
		hosts[node.Key] = host
	}
	return hosts
}
