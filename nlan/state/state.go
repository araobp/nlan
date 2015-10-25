package state

import "github.com/araobp/go-nlan/nlan/model/nlan"

type State struct {
	Router string
	Model  nlan.Model
}

type NetworkState struct {
	States []State
}
