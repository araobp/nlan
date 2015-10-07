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
	cont := context.Backgroud()
	c, _ := client.New(config)
	kapi := client.NewKeysAPI(c)

	// Key to be watched
	key := os.Args[1]

	// Creates a watcher and sets the watch
	option := client.WatcherOptions{Recursive: true}
	watcher := kapi.Watcher(key, &option)
	fmt.Printf("Watching key: %s\n", key)
	response, _ := watcher.Next(cont) // Blocks here

	// Receives a watch event
	fmt.Println("Watch event received")
	fmt.Printf("Action: %s\n", response.Action)
	node := response.Node
	fmt.Printf("Node: %s\n", node)
	fmt.Printf("Node.Key: %s\n", node.Key)
	fmt.Printf("Node.Value: %s\n", node.Value)
	fmt.Printf("Node.Nodes: %s\n", node.Nodes)
}
