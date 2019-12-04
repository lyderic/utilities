package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	network := flag.String("network", "192.168.1", "network to ping")
	flag.Parse()
	p := fastping.NewPinger()
	for i := 1; i < 255; i++ {
		ip := fmt.Sprintf("%s.%d", *network, i)
		ra, err := net.ResolveIPAddr("ip4:icmp", ip)
		if err != nil {
			log.Fatal(err)
		}
		p.AddIPAddr(ra)
	}
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %-14.14s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
