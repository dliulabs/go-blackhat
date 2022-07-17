package proxy

import (
	"io"
	"log"
	"net"
	"sync"
)

func StartProxyServer(wg *sync.WaitGroup, proxyServer net.Listener, serverAddr string) {
	defer wg.Done()

	for {
		conn, err := proxyServer.Accept()
		if err != nil {
			return
		}

		go func(from net.Conn) {
			defer from.Close()
			to, err := net.Dial("tcp", serverAddr)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer to.Close()

			err = proxy(from, to)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
		}(conn)
	}
}

func proxy(from io.Reader, to io.Writer) error {
	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if toIsReader && fromIsWriter {
		// Send replies since "from" and "to" implement the
		// necessary interfaces.
		go func() {
			_, _ = io.Copy(fromWriter, toReader)
		}()
	}

	_, err := io.Copy(to, from)
	return err
}
