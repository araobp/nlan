package util

import (
	"log"
	"testing"

	"github.com/araobp/nlan/model/nlan"
)

func TestRegisterAndListState(t *testing.T) {
	RegisterHost()
	hosts := ListHosts(true)
	log.Printf("%v", hosts)
	hosts = ListHosts(false)
	log.Printf("%v", hosts)
}

func TestSetAndGetState(t *testing.T) {
	props := nlan.VhostProps{
		Network: "10.10.10.10/24",
		Vhosts:  2,
	}
	vhosts := nlan.Vhosts{
		VhostProps: []*nlan.VhostProps{&props},
	}
	SetState("rrr", &vhosts)
	pb := new(nlan.Vhosts)
	GetState("rrr", pb)
	log.Print(pb)
	log.Print(pb.VhostProps)
}

func TestReset(t *testing.T) {
	ResetState()
}
