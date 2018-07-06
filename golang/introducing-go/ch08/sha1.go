package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("%x\n", sha1.Sum(data))
	data1 := []byte("This page intentionally left blank.1")
	fmt.Printf("%x\n", sha1.Sum(data1))

	// another approach
	h := sha1.New()
	h.Write(data)
	// The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
