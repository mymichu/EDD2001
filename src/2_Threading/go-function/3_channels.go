package main

import (
	"fmt"
	"time"
)

func source(stream chan int) {
	for i := 0; i < 10; i++ {
		stream <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(stream)
}

func main() {
	firstStream := make(chan int)

	go source(firstStream)

	for i := range firstStream {
		fmt.Println("First - Stream: ", i)
	}
	fmt.Println("Finished")
}
