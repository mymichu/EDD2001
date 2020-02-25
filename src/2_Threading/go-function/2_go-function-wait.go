package main

import (
	"fmt"
	"sync"
	"time"
)

func helloWait(wg *sync.WaitGroup, name string, number uint8) {
	defer wg.Done()
	for i := uint8(0); i < number; i++ {
		fmt.Println(i, ": Hello ", name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloWait(&wg, "Michel", 10)
	go helloWait(&wg, "ERNI", 10)
	wg.Wait()
}
