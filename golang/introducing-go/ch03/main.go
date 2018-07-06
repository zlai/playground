package main
import "fmt"

func main() {
	var x string
	x = "Hello, World"
	fmt.Println(x, x == "Hello, World")
	x += " HAHA!"
	fmt.Println(x, x == "Hello, World")
	const y string = "TRYTRY"
	fmt.Println(y)
	var (
		a = "haha"
		b = 12
		c = true
	)
	fmt.Println(a, b, c)

	var input float64
	fmt.Println("Enter a number:")
	fmt.Scanf("%f", &input)
	output := input * 10
	fmt.Println(output)
}
