package main
import "fmt"

// 2
func half(x int) (r int, even bool) {
	r = x / 2
	even = (x % 2) == 0
	return
}

// 3
func greatest(args ...int) int {
	max := 0
	for _, v := range args {
		if (v > max) {
			max = v
		}
	}
	return max
}

// 4
func makeOddGenerator() func() uint {
	x := uint(1)
	return func() (r uint) {
		r = x
		x += 2
		return
	}
}

// 5
func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

// 10
func square(x *float64) {
    *x = *x * *x
}

// 11
func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(greatest(23, 223, 4, 25,23, 1,343, 45, 4))
	odd := makeOddGenerator()
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(fib(10))
	a := 1.5
	square(&a)
	fmt.Println(a)
	x, y := 1, 20
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
