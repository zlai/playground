package main
import "fmt"

func main() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	idx := 0
	min := x[idx]
	for i, v := range x {
		if (v < min) {
			idx, min = i, v
		}
	}
	fmt.Printf("Min value: x[%d] = %d\n", idx, min)
}
