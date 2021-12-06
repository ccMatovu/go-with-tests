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
	t.Run("test the area Rectangle",func(t *testing.T) {
		rectangle := Rectangle{
						Width: 12.0,
						Height: 6.0,
					}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f but expected %.2f",got,want)
		}
	})
	
	t.Run("testing circle area",func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
        want := 314.1592653589793

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }

	})
}