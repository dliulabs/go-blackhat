package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"udpecho/monitor"
)

var (
	m  *monitor.Monitor = &monitor.Monitor{Logger: log.New(os.Stdout, "SERVER: ", 0)}
	m2 *monitor.Monitor = &monitor.Monitor{Logger: log.New(os.Stdout, "CLIENT: ", 0)}
)

func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) {
	s, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
		}()
		buf := make([]byte, 1024)
		// echo back
		for {
			n, clientAddr, err := s.ReadFrom(buf) // client to server
			if err != nil {
				return
			}
			m.Printf("Received %s from %s\n", buf[:n], clientAddr)
			_, err = s.WriteTo(buf[:n], clientAddr) // server to client
			if err != nil {
				return
			}
			m.Printf("Echoed %s\n", buf[:n])
		}
	}()
	return s.LocalAddr(), nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	m.Printf("Listening at %s\n", serverAddr)

	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = client.Close() }()
	m2.Printf("Listening at %s\n", client.LocalAddr().String())

	msg := []byte("ping")
	n, err := client.WriteTo(msg, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	m2.Printf("Send %s\n", msg[:n])

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	m2.Printf("Received %s from %s\n", buf[:n], addr)

	if addr.String() != serverAddr.String() {
		log.Fatalf("received reply from %q instead of %q", addr, serverAddr)
	}
	if !bytes.Equal(msg, buf[:n]) {
		log.Fatalf("expected reply %q; actual reply %q", msg, buf[:n])
	}
}
