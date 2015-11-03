package common

import (
	"fmt"
	"io/ioutil"
	"log"

	st "github.com/araobp/go-nlan/nlan/state"
	"github.com/araobp/go-nlan/nlan/util"
)

func ReadState(filename *string) (*[]st.State, *map[string]interface{}) {

	// Reads the state file
	state, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(state))

	// Converts YAML to Go struct
	var hosts map[string]interface{} = util.ListHosts()
	fmt.Println(hosts)
	state_ := st.NetworkState{}
	statestring := string(state)
	util.Yaml2Struct(&statestring, &state_, hosts)
	fmt.Println(state_)

	return &state_.States, &hosts
}
