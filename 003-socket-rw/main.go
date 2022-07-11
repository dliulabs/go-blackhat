package main

import (
	"io"
	"log"
	"net"
	"os"
)

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 4096)
	for {
		conn.Write([]byte("Say something ..., and I will echo back ... \n"))
		size, err := conn.Read(b[0:])
		if err != nil && err != io.EOF {
			log.Println("Unexpected error")
			panic(err)
		}

		if err == io.EOF {
			log.Println("Client disconnected")
			os.Exit(0)
		}

		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		panic(err)
	}
	log.Println("Listening on port 20080 ...")
	for {
		conn, err := listener.Accept() // blocked until accepted
		if err != nil {
			panic(err)
		}
		log.Println("Recieved connection")
		go echo(conn)
	}
}
