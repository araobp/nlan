package main

import (
	"fmt"

	br "github.com/araobp/golan/yang-test/model/bridges"
	sn "github.com/araobp/golan/yang-test/model/subnets"
	vx "github.com/araobp/golan/yang-test/model/vxlan"
)

func main() {

	enabled := true
	bridges := br.Bridges{OvsBridges: &enabled}
	fmt.Println(bridges.String())
}
