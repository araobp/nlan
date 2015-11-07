package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type resp struct {
	Error  interface{} `json:"error"`
	Id     int         `json:"id"`
	Result []rows
}

type rows struct {
	Rows []map[string]interface{} `json:"rows"`
}

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])
	var r resp
	json.Unmarshal(data, &r)
	row := r.Result[0].Rows[0]
	fmt.Print(row["ofport"])
}
