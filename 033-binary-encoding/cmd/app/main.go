package main

import (
	"binaryencoding/payload"
	"binaryencoding/types"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"reflect"
)

func Decode(r io.Reader) (types.Payload, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}

	var data types.Payload

	switch typ {
	case types.BinaryType:
		data = payload.NewBinary() // new(payload.Binary)
	case types.StringType:
		data = payload.NewString() // new(payload.String)
	default:
		return nil, errors.New("unknown type")
	}

	_, err = data.ReadFrom(
		io.MultiReader(bytes.NewReader([]byte{typ}), r))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	b1 := payload.Binary("Clear is better than clever.")
	b2 := payload.Binary("Don't panic.")
	s1 := payload.String("Errors are values.")
	payloads := []types.Payload{&b1, &s1, &b2}

	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer conn.Close()

		for _, p := range payloads {
			_, err = p.WriteTo(conn)
			if err != nil {
				log.Fatal(err)
				break
			}
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < len(payloads); i++ {
		actual, err := Decode(conn)
		if err != nil {
			log.Fatal(err)
		}

		if expected := payloads[i]; !reflect.DeepEqual(expected, actual) {
			log.Fatalf("value mismatch: %v != %v", expected, actual)
			continue
		}

		log.Printf("[%T] %[1]q", actual)
	}
}
