package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func echo(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {

		s, err := reader.ReadString('\n') // read up to the delimiter \n character.
		if err == io.EOF {
			log.Println("Disconnected.")
			os.Exit(0)
		} else if err != nil {
			switch v := err.(type) {
			default:
				log.Fatalf("Unable to read data %v", v)
			}
		}
		log.Printf("Read %d bytes of data: %s\n", len(s), s)
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		writer.Flush()
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
