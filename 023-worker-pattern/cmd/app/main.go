package main

import "fmt"

func worker(track chan bool, task chan int) {
	for i := range task {
		fmt.Printf("work: %d\n", i)
	}
	track <- true
}

func main() {
	task := make(chan int)
	track := make(chan bool)
	defer close(track)     // close track channel
	go worker(track, task) // msut start worker first
	for i := 1; i < 10; i++ {
		task <- i
	}
	close(task) // tells workers, no more task
	<-track     // wait for worker to complete
}
