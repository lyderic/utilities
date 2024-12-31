package main

import (
	"fmt"
	"net"
)

func main() {
	iface, _ := net.InterfaceByName("tailscale0")
	addrs, _ := iface.Addrs()
	for _, addr := range addrs {
		if ip := addr.(*net.IPNet).IP.To4(); ip != nil {
			fmt.Println(ip.String())
			break
		}
	}
}
