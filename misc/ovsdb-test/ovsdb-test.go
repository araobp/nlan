// 2015/10/15
// Minimal RFC7047(OVSDB mgmt protocol) implementation for one-shot operation.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

const (
	DATABASE = "Open_vSwitch"
	SOCK     = "/var/run/openvswitch/db.sock"
)

type jsonrpc struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	Id     int           `json:"id"`
}

type Operation struct {
	Op    string        `json:"op"`
	Table string        `json:"table"`
	Where []interface{} `json:"where"`
}

func Condition(column string, function string, value interface{}) []interface{} {
	return []interface{}{column, function, value}
}

func read(conn net.Conn) []byte {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := buf[0:n]
	return response
}

var id int = 0

// Synchronous JSON-RPC request to OVSDB.
// Note: root privilege required for executing this API.
func RequestSync(method string, params []interface{}) []byte {
	conn, err := net.Dial("unix", SOCK)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	id++
	rpc := jsonrpc{
		Method: method,
		Params: params,
		Id:     id,
	}
	json_data, err := json.Marshal(rpc)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Request: %s\n", string(json_data))
	fmt.Fprintf(conn, string(json_data))
	response := read(conn)
	log.Printf("Response: %s\n", string(response))
	return response
}

func main() {
	_ = RequestSync("list_dbs", []interface{}{})

	cond := Condition("vni", "==", 1001)
	ope := Operation{
		Op:    "select",
		Table: "Interface",
		Where: cond,
	}
	_ = RequestSync("transact", []interface{}{DATABASE, ope})
}
