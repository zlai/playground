package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"errors"
)

func main() {
	f, err := os.Open("main.go")
	if err != nil {
		return
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = f.Read(bs)
	if err != nil {
		return
	}
	// convert to string for printing
	fmt.Println(string(bs))


	// using ioutil
	bs, err = ioutil.ReadFile("main.go")
	fmt.Println(string(bs))

	// dir
	dir, err := os.Open(".")
	defer dir.Close()
	info, err := dir.Readdir(-1)
	for _, f := range info {
		fmt.Println(f.Name())
	}

	e := errors.New("A TEST ERROR!")
	fmt.Println(e)
}
