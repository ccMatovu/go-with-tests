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
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
		{Triangle{12,6},36.0},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}