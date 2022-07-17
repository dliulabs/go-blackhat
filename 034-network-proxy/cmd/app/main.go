package main

import (
	"log"
	"net"
	"sync"

	"proxycopy/proxy"
	"proxycopy/server"
)

const (
	SERVERADDR  = "localhost:9090"
	PROXYSERVER = "localhost:8080"
)

func main() {
	var wg sync.WaitGroup
	svr, err := net.Listen("tcp", SERVERADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer svr.Close()

	wg.Add(1)
	go server.StartServer(&wg, svr)

	proxyServer, err := net.Listen("tcp", PROXYSERVER)
	if err != nil {
		log.Fatal(err)
	}
	defer proxyServer.Close()

	wg.Add(1)
	go proxy.StartProxyServer(&wg, proxyServer, SERVERADDR)

	conn, err := net.Dial("tcp", proxyServer.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	msgs := []struct{ Message, Reply string }{
		{"ping", "pong"},
		{"pong", "pong"},
		{"echo", "echo"},
		{"ping", "pong"},
	}

	for i, m := range msgs {
		_, err = conn.Write([]byte(m.Message))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%d: expected reply: %q\n", i, m.Reply)
	}

	wg.Wait()
}
