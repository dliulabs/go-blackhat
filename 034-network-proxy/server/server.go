package server

import (
	"io"
	"log"
	"net"
	"sync"
)

func StartServer(wg *sync.WaitGroup, server net.Listener) {
	defer wg.Done()

	for {
		conn, err := server.Accept()
		if err != nil {
			return
		}

		go func(c net.Conn) {
			defer c.Close()

			for {
				buf := make([]byte, 1024)
				n, err := c.Read(buf)
				if err != nil {
					if err != io.EOF {
						log.Fatal(err)
					}

					return
				}
				log.Printf("SERVER: received %d bytes, %s\n", n, string(buf[:n]))
				switch msg := string(buf[:n]); msg {
				case "ping":
					_, err = c.Write([]byte("pong"))
				default:
					_, err = c.Write(buf[:n])
					log.Printf("SERVER: echo %d bytes %s\n", n, buf[:n])
				}

				if err != nil {
					if err != io.EOF {
						log.Fatal(err)
					}

					return
				}
			}
		}(conn)
	}

}
