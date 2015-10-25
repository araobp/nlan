package main

import (
	"fmt"

	"github.com/araobp/go-nlan/misc/grpc-test/api"
	"github.com/golang/protobuf/proto"
)

func main() {

	bridges := api.Bridges{Dummy: "test"}
	data, _ := proto.Marshal(&bridges)
	fmt.Println(data)

	bridges_ := api.Bridges{}
	_ = proto.Unmarshal(data, &bridges_)
	fmt.Println(bridges_)

}
