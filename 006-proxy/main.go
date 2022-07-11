package main

import (
	"fmt"
	"io"
	"net"
)

// copy source to dest

var localaddr = "localhost:9090"
var remoteaddr = "tcpbin.com:4242" // see https://tcpbin.com/

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", remoteaddr)
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	fmt.Printf("Proxy to %s\n", remoteaddr)
	src.Write([]byte(fmt.Sprintf("Proxy connected to %s\n", remoteaddr)))
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			panic(err)
		}

	}()
	if _, err := io.Copy(src, dst); err != nil {
		panic(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", localaddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening at %s\n", localaddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		go handle(conn)

	}
}
