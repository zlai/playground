package main
import "fmt"

func main() {
	var f float64
	fmt.Println("Enter Fahrenheit: ")
	fmt.Scanf("%f", &f)
	fmt.Println("Celsius: ", (f-32)*5/9)
}
