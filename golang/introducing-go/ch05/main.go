package main
import "fmt"

func main() {
	x := [5]float64 {
		98,
		93,
		77,
		82,
		83,
	}
	sum := 0.0
	for _, v := range x {
		sum += v
	}
	fmt.Println("Average: ", sum/float64(len(x)))


	// slice append
	x1 := []int {1, 2, 3}
	x2 := append(x1, 4, 5)
	fmt.Println(x1, x2)

	// slice copy
	a1 := []int {2, 4, 5}
	a2 := make([]int, 2)
	copy(a2, a1)
	fmt.Println(a1, a2)

	// map
	y := make(map[string]int)
	y["world"] = 2
	fmt.Println(y["world"])

	// map create
	ele := map[string]string {
		"a" : "AA",
		"b" : "BBB",
		"c" : "CCCC",
	}
	fmt.Println(ele)
	if val, ok := ele["c"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Not found!")
	}
}
