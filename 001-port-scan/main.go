package main

import (
	"fmt"
	"net"
	"sync"
)

type Result struct {
	addr string
	open bool
	err  error
}

var results = make(map[string]Result)

func Scan(addr string, m *sync.Mutex) {
	_, err := net.Dial("tcp", addr)
	m.Lock()
	defer m.Unlock()
	results[addr] = Result{
		addr: addr,
		open: err == nil,
		err:  err,
	}
	return
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	host := "scanme.nmap.org"
	for port := 20; port < 65535; port++ {
		wg.Add(1)
		addr := fmt.Sprintf("%s:%d", host, port)
		go func(addr string) {
			defer wg.Done()
			Scan(addr, &mu)
		}(addr)
	}
	wg.Wait()
	for addr, result := range results {
		if result.open {
			fmt.Printf("%s: %v\n", addr, result.open)
		}
	}
}
