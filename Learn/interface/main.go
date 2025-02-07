package main

import (
	"fmt"
)

// Định nghĩa interface Shape
type Shape interface {
	Area() float64
}

// Struct Circle
type Circle struct {
	Radius float64
}

// Struct Rectangle
type Rectangle struct {
	Width, Height float64
}

// Implement phương thức Area() cho Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement phương thức Area() cho Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Khai báo một biến kiểu Shape
	var s Shape

	// Gán một Circle vào s
	s = Circle{}
	fmt.Println("Circle Area:", s.Area())

	// Gán một Rectangle vào s
	s = Rectangle{Width: 4, Height: 6}
	fmt.Println("Rectangle Area:", s.Area())
}
