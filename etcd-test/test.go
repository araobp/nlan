package main

import (
	"log"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

func main() {

	config := client.Config{
		Endpoints:               []string{"http://localhost:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: Time.Second,
	}

	c, err := client.New(config)

	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Set(context.Background(), "foo", "bar", nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(resp)
	}
}
