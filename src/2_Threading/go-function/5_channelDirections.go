package main

import (
	pingpong "EDD/src/2_Threading/go-function/5_pingpong"
	"fmt"
)

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	pingpong.Ping(pings, "passed message")
	pingpong.Pong(pings, pongs)
	fmt.Println(<-pongs)
}
