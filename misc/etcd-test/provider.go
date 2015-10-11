package main

import (
	"fmt"
	"os"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

func main() {

	// etcd config setup
	config := client.Config{
		Endpoints:               []string{"http://localhost:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	// Creates an etcd client
	cont := context.Background()
	c, _ := client.New(config)
	kapi := client.NewKeysAPI(c)

	// Key and value
	key := "sdn2"
	value := os.Args[1]

	// Appends a value to a list on etcd
	_, _ = kapi.CreateInOrder(cont, key, value, nil)

	response, _ := kapi.Get(cont, key, &client.GetOptions{Recursive: true})
	nodes := response.Node.Nodes
	// Dumps values in a list on etcd
	for _, v := range nodes {
		fmt.Println(v.Value)
	}
}
