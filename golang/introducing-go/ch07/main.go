package main
import ("fmt"; "math")

type Circle struct {
	r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 1
}

func (r *Rectangle) area() float64 {
	return 1
}

func (r *Rectangle) perimeter() float64 {
	return 1
}

func totalArea(shapes ...Shape) float64 {
	a := 0.0
	for _, s := range shapes {
		a += s.area()
	}
	return a
}
func main() {
	//var c Circle
	//c := new(Circle)
	c := Circle{1, 2, 4}
	fmt.Println(c)
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())
}
