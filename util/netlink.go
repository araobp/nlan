package util 

import (
	"github.com/vishvananda/netlink"

	"net"
)

type Route struct {
	Via net.IP
	Dev string
	Src net.IP
}

// Returns a list of devices and addresses ("ip addr show")
func AddrMap() (*map[string][]string, *map[string]string) {
	addrs, _ := netlink.AddrList(nil, netlink.FAMILY_V4)
	devMap := make(map[string][]string)
	addrMap := make(map[string]string)
	for _, a := range addrs {
		addr := a.IP.String()
		dev := a.Label
		l, ok := devMap[dev]
		if ok {
			devMap[dev] = append(l, addr)
		} else {
			devMap[dev] = []string{addr}
		}
	}
	return &devMap, &addrMap
}

// Returns a list of routes
func RouteMap() *map[string]Route {
	links, _ := netlink.LinkList()
	linksMap := make(map[int]string)
	for _, l := range links {
		attrs := *l.Attrs()
		linksMap[attrs.Index] = attrs.Name
	}

	routes := make(map[string]Route)
	routeList, _ := netlink.RouteList(nil, netlink.FAMILY_V4)
	for _, r := range routeList {
		if_ := linksMap[r.LinkIndex]
		rdst := r.Dst
		var dst string
		if rdst != nil {
			dst = rdst.String()
		} else {
			dst = "default"
		}

		route := Route{
			Via: r.Gw,
			Dev: if_,
			Src: r.Src,
		}
		routes[dst] = route
	}
	return &routes
}
		
