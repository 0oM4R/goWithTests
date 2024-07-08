package perimeter

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
func (c Circle) Area() float64 { return math.Pi * math.Pow(c.Radius, 2) }
