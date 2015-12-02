// 2015/10/15
// Minimal RFC7047(OVSDB mgmt protocol) implementation for one-shot operation.
package util

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

const (
	DATABASE = "Open_vSwitch"
	SOCK     = "/var/run/openvswitch/db.sock"
	BUFSIZE  = 4096
)

type jsonrpc struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	Id     int           `json:"id"`
}

type Operation struct {
	Op      string        `json:"op"`
	Table   string        `json:"table"`
	Where   []interface{} `json:"where"`
	Columns []string      `json:"columns"`
}

type OperationWoColumns struct {
	Op    string        `json:"op"`
	Table string        `json:"table"`
	Where []interface{} `json:"where"`
}

type Response struct {
	Error  interface{} `json:"error"`
	Id     int         `json:"id"`
	Result []Rows
}

type Rows struct {
	Rows []map[string]interface{} `json:"rows"`
}

func Condition(column string, function string, value interface{}) []interface{} {
	return []interface{}{column, function, value}
}

func read(conn net.Conn) []byte {
	buf := make([]byte, BUFSIZE)
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

// Fetches ofport from OVSDB.
func GetOfport(port string) int {
	cond := Condition("name", "==", port)
	ope := OperationWoColumns{
		Op:    "select",
		Table: "Interface",
		Where: []interface{}{cond},
	}
	resp := RequestSync("transact", []interface{}{DATABASE, ope})

	var r Response
	json.Unmarshal(resp, &r)
	row := r.Result[0].Rows[0]
	ofport := int(row["ofport"].(float64))
	log.Printf("ofport: %d\n", ofport)
	return ofport
}

// Fetches VXLAN ofport from OVSDB.
func GetVxlanPorts(peers *[]string) *list.List {
	cond := Condition("type", "==", "vxlan")
	ope := Operation{
		Op:      "select",
		Table:   "Interface",
		Where:   []interface{}{cond},
		Columns: []string{"ofport", "options"},
	}
	resp := RequestSync("transact", []interface{}{DATABASE, ope})
	var r Response
	json.Unmarshal(resp, &r)
	rows := r.Result[0].Rows
	l := list.New()
	for _, ip := range *peers {
		for _, row := range rows {
			options := row["options"].([]interface{})
			for _, e := range options[1].([]interface{}) {
				var ip_ string
				elm := e.([]interface{})
				if elm[0].(string) == "remote_ip" {
					ip_ = elm[1].(string)
					if ip == ip_ {
						ofport := int(row["ofport"].(float64))
						log.Printf("ofport: %d\n", ofport)
						l.PushBack(ofport)
					}
					break
				}
			}
		}
	}
	return l
}

func GetBridgeNames() *list.List {
	ope := Operation{
		Op:      "select",
		Table:   "Bridge",
		Where:   []interface{}{}, // empty
		Columns: []string{"name"},
	}
	resp := RequestSync("transact", []interface{}{DATABASE, ope})
	var r Response
	json.Unmarshal(resp, &r)
	rows := r.Result[0].Rows
	l := list.New()
	for _, row := range rows {
		bridge := row["name"].(string)
		l.PushBack(bridge)
	}
	return l
}
