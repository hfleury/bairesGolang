package shapes

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height

}

func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func calculate(s Shape) {
	fmt.Printf("The area is: %f \n", s.Area())
	fmt.Printf("The perimeter is: %f \n", s.Perimeter())
}

// Shapesmain is the main function of this file
func Shapesmain() {
	rectangle := Rectangle{10, 20}
	circle := Circle{43}
	calculate(rectangle)
	calculate(circle)
}
