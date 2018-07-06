package main

import (
	"fmt"
	"time"
)

func MySleep(n time.Duration) {
	select {
	case <-time.After(n):
		return
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			MySleep(100)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			MySleep(200)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(80):
				fmt.Println("TIMEOUT!")
			default:
				fmt.Println("nothing to process!")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
