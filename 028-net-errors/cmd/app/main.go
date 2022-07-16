package main

import (
	"log"
	"net"
	"syscall"
	"time"
)

func DialTimeout(network, address string, timeout time.Duration,
) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func main() {
	c, err := DialTimeout("tcp", "localhost:8080", 5*time.Second)
	if err == nil {
		c.Close()
		log.Fatal("connection did not time out")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		log.Fatal(err)
	}
	if !nErr.Timeout() {
		log.Fatal("error is not a timeout")
	}
	if nErr.Timeout() {
		log.Fatalf("Error: %s", err.Error())
	}
}
