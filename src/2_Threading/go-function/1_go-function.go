package main

import (
	"fmt"
	"time"
)

func hello(name string, number uint8) {
	for i := uint8(0); i < number; i++ {
		fmt.Println(i, ": Hello ", name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go hello("Michel", 10)
}
