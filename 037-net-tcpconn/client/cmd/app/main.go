package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const defaultPingInterval = 3 * time.Second

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":9000")
	if err != nil {
		log.Fatalf("ResolveTCPAddr returned: %s", err)
	}
	// client, err := net.Dial("tcp", "localhost:9090")
	client, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// subscribing to os.Signal(1), or SIGHUP (hang up)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// a ticker is a channel that produces a message at each tick
	// NewTicker returns a new Ticker containing a channel that will send the current time on the channel after each tick
	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()
	timer := time.NewTimer(defaultPingInterval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	b := make([]byte, 4096)
	for {
		select {
		// case <-ticker.C:
		case <-timer.C:
			// relay every tick to the server
			_, err := client.Write([]byte("ping"))
			if err != nil {
				log.Fatal("write:", err)
			}

			err = client.SetReadDeadline(time.Now().Add(1 * time.Second))
			if err != nil {
				panic(err)
			}
			n, err := client.Read(b[0:])
			if err != nil && err != io.EOF {
				log.Println("Unexpected error")
				panic(err)
			}

			if err == io.EOF {
				log.Println("Client disconnected")
				os.Exit(0)
			}

			log.Printf("Received %d bytes: %s", n, string(b))
			if string(b[:n]) != "PONG!!!" {
				log.Fatalln("failed to receive PONG!!!")
			}
			_ = timer.Reset(defaultPingInterval) // must reset otherwise it will stop

		case <-interrupt:
			// SIGHUP
			log.Println("interrupt")
			os.Exit(0)

		}
	}
}
