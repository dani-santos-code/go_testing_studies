package structs

import "math"

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a rectangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Shape interface
type Shape interface {
	Area() float64
}

// Perimeter calculates the perimeter of a rectangle
// func Perimeter(r Rectangle) float64 {
// 	return 2 * (r.Width + r.Height)
// }
