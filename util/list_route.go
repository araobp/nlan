package util 

import (
	"github.com/vishvananda/netlink"

	"fmt"
)

type Route struct {
	*netlink.Route
	If     string
	IfAddr string
}

func RouteList() []Route {
	addrs, _ := netlink.AddrList(nil, netlink.FAMILY_V4)
	addrsMap := make(map[string]string)
	for _, a := range addrs {
		addrsMap[a.Label] = a.IP.String()
	}
	links, _ := netlink.LinkList()
	linksMap := make(map[int]string)
	for _, l := range links {
		attrs := *l.Attrs()
		linksMap[attrs.Index] = attrs.Name
	}
	routeList, _ := netlink.RouteList(nil, netlink.FAMILY_V4)
	var routes []Route

	for _, r := range routeList {
		if_ := linksMap[r.LinkIndex]
		ifAddr := addrsMap[if_]
		fmt.Println(r)
		route := Route{Route: &r,
			If:     if_,
			IfAddr: ifAddr,
		}
		routes = append(routes, route)
	}
	return routes
}
