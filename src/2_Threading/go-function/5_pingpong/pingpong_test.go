package pingpong

import (
	"testing"
)

func TestPing(t *testing.T) {
	pings := make(chan string, 1)
	Ping(pings, "hello")
	result := <-pings
	if result != "hello" {
		t.Errorf("expected hello - received %s", result)
	}
}

func TestPong(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	Ping(pings, "hello")
	Pong(pings, pongs)
	result := <-pongs
	if result != "hello" {
		t.Errorf("expected hello - received %s", result)
	}
}

func BenchmarkAddition(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Addition(10, 30)
	}
}
