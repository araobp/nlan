package ptn

import (
	"log"

	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
)

func Add(in *nlan.Ptn) {
	networks := in.GetNetworks()
	for _, net := range networks {
		log.Print(net)
	}
}

func Update(in *nlan.Ptn) {
	log.Print(in)
}

func Delete(in *nlan.Ptn) {
	log.Print(in)
}
