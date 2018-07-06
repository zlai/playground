package main
import "fmt"

func average(xs []float64) (a float64) {
	sum := 0.0
	for _, v := range xs {
		sum += v
	}
	a = sum / float64(len(xs))
	return
}

// variadic parameter
func sum(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

// closure
func makeEven() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// panic, recover
func testPanic() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("JUST PANIC")
}

// pointer
func reset(v *int) {
	*v = 100
}

func main() {
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
	y := []int{98,93,77,82,83}
	fmt.Println(sum(y...))
	efun := makeEven()
	fmt.Println(efun())
	fmt.Println(efun())
	fmt.Println(efun())
	fmt.Println(efun())
	testPanic()
	z := 800
	fmt.Println(z)
	reset(&z)
	fmt.Println(z)

	k := new(int)
	fmt.Println(*k)
	reset(k)
	fmt.Println(*k)
}

