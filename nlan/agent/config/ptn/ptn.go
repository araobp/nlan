package ptn

import (
	"log"

	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
)

func Add(in *nlan.Ptn, logger *log.Logger) {
	networks := in.GetNetworks()
	for _, net := range networks {
		logger.Print(net)
	}
}

func Update(in *nlan.Ptn, logger *log.Logger) {
	logger.Print(in)
}

func Delete(in *nlan.Ptn, logger *log.Logger) {
	logger.Print(in)
}
