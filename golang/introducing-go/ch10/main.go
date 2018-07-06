package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(200))
		time.Sleep(amt * time.Millisecond)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
