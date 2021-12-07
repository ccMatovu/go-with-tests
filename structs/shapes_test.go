package structs

import (
	"testing"
)

type Rectangle struct{
	Width float64
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