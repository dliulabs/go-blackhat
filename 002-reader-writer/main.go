package main

import (
	"fmt"
	"os"
)

type FooReader struct{}

type FooWriter struct{}

func (r *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

func (w *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out >")
	return os.Stdout.Write(b)
}

func main() {

	var (
		reader FooReader
		writer FooWriter
	)
	var input = make([]byte, 4096)
	count, err := reader.Read(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes\n", count)

	count, err = writer.Write(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", count)
}
