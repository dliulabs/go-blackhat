package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// Create a listener
	// l, err := net.Listen("tcp", ":9000")
	addr, err := net.ResolveTCPAddr("tcp", ":9000")
	if err != nil {
		log.Fatalf("ResolveTCPAddr returned: %s", err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Listener returned: %s", err)
	}
	defer l.Close()

	for {
		// Accept new connections
		// c, err := l.Accept()
		c, err := l.AcceptTCP()
		if err != nil {
			log.Fatalf("Unable to accept new connections: %s", err)
		}

		// Create a goroutine that reads and writes-back data
		go func() {
			log.Printf("TCP Session Open")
			// Clean up session when goroutine completes, it's ok to call Close more than once.
			defer c.Close()

			for {
				b := make([]byte, 120)

				// Read from TCP Buffer
				n, err := c.Read(b)
				if err != nil {
					log.Printf("Error reading TCP Session: %s", err)
					break
				}
				log.Printf("Received %d bytes %s\n", n, b[:n])
				// Write-back data to TCP Client
				if string(b[:n]) == "ping" {
					_, err = c.Write([]byte("PONG!!!"))
				} else {
					_, err = c.Write(b)
				}
				if err != nil {
					log.Printf("Error writing TCP Session: %s", err)
					break
				}
			}
		}()

		// Create a goroutine that closes a session after 30 seconds
		go func() {
			<-time.After(time.Duration(30) * time.Second)

			// Use SetLinger to force close the connection
			// immediately discard unsent data on close
			// err := c.(*net.TCPConn).SetLinger(0)
			err := c.SetLinger(0)
			if err != nil {
				log.Printf("Error when setting linger: %s", err)
			}

			defer c.Close()
		}()
	}
}
