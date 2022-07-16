package main

import (
	"fmt"

	"inet.af/netaddr"
)

func main() {
	var ipaddr = netaddr.IPv4(192, 168, 1, 10)
	fmt.Println(ipaddr)

	var ip2 netaddr.IP
	if err := ip2.UnmarshalBinary([]byte{192, 168, 1, 10}); err == nil {
		fmt.Printf("unmarshaled into non-empty IP: %s\n", ip2)
	}

	var prefix netaddr.IPPrefix = netaddr.IPPrefixFrom(netaddr.IPv4(192, 168, 156, 97), 19)
	fmt.Printf("Prefix Bits: %d\n", prefix.Bits())
	fmt.Printf("Prefix IP: %s\n", prefix.IP())
	fmt.Printf("Prefix IPNet Mask: %s\n", prefix.IPNet().Mask)
}
