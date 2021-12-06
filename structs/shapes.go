package structs

import "math"



func  Perimeter(rectangle Rectangle)float64{
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area()float64{
	return (r.Width * r.Height)
}

func (c Circle) Area()float64{
	return math.Pi * c.Radius * c.Radius
}