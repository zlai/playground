package main
import (
	"fmt"
	"time"
)

func pinger(ch chan<- string) {
	for {
		ch <- "ping"
	}
}

func ponger(ch chan<- string) {
	for {
		ch <- "pong"
	}
}

func printer(ch <-chan string) {
	for {
		fmt.Println(<- ch)
		time.Sleep(5)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

