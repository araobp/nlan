package common

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	st "github.com/araobp/nlan/state"
	"github.com/araobp/nlan/util"
)

const LOGDIR = "/var/volume/"

// This function reads NLAN state file from a specified path.
// Set roster to nil if roster is on etcd.
func ReadState(filename *string, roster *string) (*[]st.State, *map[string]interface{}) {

	// Reads the state file
	state, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(state))

	// Converts YAML to Go struct
	var hosts map[string]interface{}
	var secondary map[string]interface{}

	if roster == nil {
		hosts = util.ListHosts(false)
		secondary = util.ListHosts(true)
	} else {
		r, err := ioutil.ReadFile(*roster)
		if err != nil {
			log.Fatal(err)
		}
		yaml.Unmarshal(r, &hosts)
	}
	log.Print(hosts)
	state_ := st.NetworkState{}
	statestring := string(state)
	if len(secondary) > 0 {
		util.Yaml2Struct(&statestring, &state_, secondary)
	} else {
		util.Yaml2Struct(&statestring, &state_, hosts)
	}
	log.Print(state_)

	return &state_.States, &hosts
}

func WriteLog(filename string, buf *bytes.Buffer) error {
	err := ioutil.WriteFile(LOGDIR+filename, buf.Bytes(), 0644)
	buf.Reset()
	return err
}
