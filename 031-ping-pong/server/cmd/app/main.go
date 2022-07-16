package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			return
		}

		log.Printf("received: %q", buf[:n])
		if string(buf[0:4]) == "ping" {
			_, err = conn.Write([]byte("PONG!!!"))
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}
	for {
		log.Println("Listening at localhost:9090")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}
