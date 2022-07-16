package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	timeout := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), timeout)
	defer cancel()

	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			// Sleep long enough to reach the context's deadline.
			time.Sleep(2*time.Second + time.Millisecond)
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
	}
	sync := make(chan struct{})

	go func() {
		defer func() { sync <- struct{}{} }()

		conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
		if err == nil {
			conn.Close()
			log.Println("connection did not time out")
		}
		nErr, ok := err.(net.Error)
		if !ok {
			log.Println(err)
		} else {
			if !nErr.Timeout() {
				log.Printf("error is not a timeout: %v", err)
			}
		}
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("expected deadline exceeded; actual: %v", ctx.Err())
			return
		}
		log.Println("connection did not time out")
	}()

	<-sync
}
