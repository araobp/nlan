package main

import (
	"encoding/json"
	"log"
	_ "net"

	"github.com/socketplane/libovsdb"
)

func main() {
	//conn, err := net.Dial("tcp", "localhost:6379")
	//if err != nil {
	//}
	type Rpc struct {
		method string
		params [][]byte
		id     int
	}
	condition := libovsdb.NewCondition("vni", "==", 1001)
	ope := libovsdb.Operation{
		Op:    "select",
		Table: "interface",
		Where: condition,
	}
	p := []libovsdb.Operation{ope}
	p0, _ := json.Marshal(p)
	log.Printf("%v\n", string(p0))
	p := Rpc{method: "transact", params: [][]byte{p0}, id: 1}
	json_data, _ := json.Marshal(p)
	log.Printf("%v\n", string(json_data))
}
