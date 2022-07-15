package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var ServerAddr = "localhost:9090"

func main() {
	listener, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// done := make(chan struct{})
	go func() {
		// defer func() { done <- struct{}{} }()

		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}

			go func(c net.Conn) {
				defer func() {
					c.Close()
					// done <- struct{}{}
				}()

				r := bufio.NewReader(c)
				time.Sleep(time.Second / 2)
				for {
					msg, err := r.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							panic(err)
						}
					}

					fmt.Printf("received: %q\n", msg)
				}
			}(conn)
		}
	}()

	client, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// subscribing to os.Signal(1), or SIGHUP (hang up)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// a ticker is a channel that produces a message at each tick
	// NewTicker returns a new Ticker containing a channel that will send the current time on the channel after each tick
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		// case <-done:
		// return
		case t := <-ticker.C:
			// relay every tick to the server
			_, err := client.Write([]byte(fmt.Sprintf("it is now %s\n", t.String())))
			if err != nil {
				log.Println("write:", err)
				return
			}

		case <-interrupt:
			// SIGHUP
			log.Println("interrupt")
			os.Exit(0)
		}
	}
}
