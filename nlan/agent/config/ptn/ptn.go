package ptn

import (
	"log"

	nlan "github.com/araobp/go-nlan/nlan/model/nlan"
)

func GetElements(in *nlan.Ptn) ([]*nlan.PtnNodes, []*nlan.PtnLinks, []*nlan.PtnL2Vpn) {
	return in.GetPtnNodes(), in.GetPtnLinks(), in.GetPtnL2Vpn()
}

func Add(in *nlan.Ptn) {
	ptn_nodes, ptn_links, ptn_l2vpn := GetElements(in)
	if ptn_nodes != nil {
		log.Print(ptn_nodes)
	}
	if ptn_links != nil {
		log.Print(ptn_links)
	}
	if ptn_l2vpn != nil {
		log.Print(ptn_l2vpn)
	}
}

func Update(in *nlan.Ptn) {
	log.Print(in)
}

func Delete(in *nlan.Ptn) {
	log.Print(in)
}
