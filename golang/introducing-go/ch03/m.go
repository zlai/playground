package main
import "fmt"

func main() {
	var f float64
	fmt.Println("Enter Feet: ")
	fmt.Scanf("%f", &f)
	fmt.Println("Meters: ", f*0.3048)
}
