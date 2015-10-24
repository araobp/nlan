package state

import "github.com/araobp/golan/nlan/model/nlan"

type State struct {
	Router string
	Model  nlan.Model
}

type NetworkState struct {
	States []State
}
