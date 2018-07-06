package main
import (
	"fmt"
	"container/list"
)

func main() {
	var l list.List
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
		fmt.Println(e.Value)
	}
}
