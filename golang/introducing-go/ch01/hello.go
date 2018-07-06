package main

import "fmt"
import "os"

func main() {
	defer func() {fmt.Println("DEFERRED!")} ()
	fmt.Println("Hello World!")
	os.Exit(100)
}
