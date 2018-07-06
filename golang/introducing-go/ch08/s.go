package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (ns ByName) Len() int {
	return len(ns)
}

func (ns ByName) Less(i, j int) bool {
	return ns[i].Name < ns[j].Name
}

func (ns ByName) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	fmt.Println(kids)
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
