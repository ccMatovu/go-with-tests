package structs

import (
	"fmt"
	"testing"
)

type Rectangle struct{
	Width float64
	Height float64
}

type Triangle struct{
	Base float64
	Height float64
}

type Circle struct{
	Radius float64
}

type Shape interface{
	Area() float64
}

func TestPerimeter(t *testing.T){
	t.Run("testing the perimeter",func(t *testing.T) {
		rectangle := Rectangle{10.0,10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f but expected %.2f", got, want)
		}
	})	
}
/*
func TestArea(t *testing.T){

	checkArea := func (t *testing.T,shape Shape,want float64)  {
		t.Helper()
		got := shape.Area()

		if got != want{
			t.Errorf("got %g but expected %g",got,want)
		}
	}

	t.Run("test the area Rectangle",func(t *testing.T) {
		rectangle := Rectangle{
						Width: 12.0,
						Height: 6.0,
					}
		want := 72.0
		checkArea(t,rectangle,want)
	})
	
	t.Run("testing circle area",func(t *testing.T) {
		circle := Circle{10}
        want := 314.1592653589793
		checkArea(t,circle,want)
	})
}

*/
func BenchmarkPerimeter(b *testing.B){
	rectangle := Rectangle{10.0,10.0}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}

func ExamplePerimeter(){
	rectangle := Rectangle{10.0,10.0}
	fmt.Println(Perimeter(rectangle))
	//Output: 40
}

func TestArea(t *testing.T) {

    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
            }
        })

	}
}