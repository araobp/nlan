package common

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	st "github.com/araobp/go-nlan/nlan/state"
	"github.com/araobp/go-nlan/nlan/util"
)

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
	if roster == nil {
		hosts = util.ListHosts()
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
	util.Yaml2Struct(&statestring, &state_, hosts)
	log.Print(state_)

	return &state_.States, &hosts
}

func WriteLog(filename string, buf *bytes.Buffer) error {
	b := *buf
	return ioutil.WriteFile("/tmp/"+filename, b.Bytes(), 0644)
}
