package pingpong

func Ping(pings chan<- string, msg string) {
	pings <- msg
}

func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func Addition(sum1 int16, sum2 int16) int16 {
	return sum1 + sum2
}
