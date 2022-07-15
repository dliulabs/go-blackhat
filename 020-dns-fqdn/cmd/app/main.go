package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	r, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		log.Fatalf("failed to exchange: %v", err)
	}
	if r == nil {
		log.Fatal("response is nil")
	}
	if r.Rcode != dns.RcodeSuccess {
		fmt.Errorf("expected rcode %v, got %v", dns.RcodeSuccess, r.Rcode)
	}
	fmt.Printf("%v\n", r)

	for _, answer := range r.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
