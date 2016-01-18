package common

import (
	"io/ioutil"
	"log"

	"github.com/araobp/nlan/model/nlan"
	"github.com/araobp/nlan/util"
)

const LOGDIR = "/var/volume/"

// This function reads NLAN state file from a specified path.
func ReadState(filename *string) (*nlan.State, *map[string]interface{}) {

	state := nlan.State{}

	// Reads the state file
	statebyte, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(statebyte))

	// Converts YAML to Go struct
	var hosts map[string]interface{}

	hosts = util.ListIps()
	log.Print(hosts)
	statestring := string(statebyte)
	util.Yaml2Struct(&statestring, &state.Router, hosts)
	log.Print(state)

	return &state, &hosts
}
