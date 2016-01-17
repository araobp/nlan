// Master
package main

import (
	"flag"
	"log"
	"os"

	"github.com/araobp/nlan/common"
	"github.com/araobp/nlan/util"
)

func main() {
	filename := flag.String("state", "state.yaml", "state file")
	reset := flag.Bool("reset", false, "reset NLAN state on tega")
	flag.Parse()

	if *reset == true {
		util.ResetState()
		os.Exit(0)
	}

	log.Println(*filename)

	states, _ := common.ReadState(filename, nil)

	// Writes state onto tega db
	for _, v := range *states {
		router := v.Router
		state := v.Model
		util.SetState(router, &state)
	}
}
