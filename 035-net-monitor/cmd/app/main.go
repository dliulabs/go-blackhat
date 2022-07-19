package main

import (
	"io"
	"log"
	"net"
	"netmonitor/monitor"
	"os"
)

func main() {
	ExampleMonitor()
}

func ExampleMonitor() {
	m := &monitor.Monitor{Logger: log.New(os.Stdout, "SERVER: ", 0)}
	m2 := &monitor.Monitor{Logger: log.New(os.Stdout, "MONITOR: ", 0)}

	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		m.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		b := make([]byte, 1024)
		r := io.TeeReader(conn, m)
		n, err := r.Read(b)
		if err != nil && err != io.EOF {
			m.Println(err) // Server
			return
		}

		w := io.MultiWriter(conn, m2)
		_, err = w.Write(b[:n]) // echo the message
		if err != nil && err != io.EOF {
			m2.Println(err) // Monitor
			return
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		m2.Fatal(err)
	}

	_, err = conn.Write([]byte("Test\n"))
	if err != nil {
		m2.Fatal(err)
	}

	_ = conn.Close()
	<-done
}
