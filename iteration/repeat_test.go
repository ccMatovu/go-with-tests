package iteration

import (
	"fmt"
	"testing"
)
func TestRepeat(t *testing.T){
	repeated := Repeat("a",4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q",expected,repeated)
	}
}

func ExampleRepeat(){
	example := Repeat("a",4)
	fmt.Println(example)
	//Output: aaaa
}

func BenchmarkRepeat(b *testing.B){
	for i:= 0; i < b.N; i++{
		Repeat("a",4)
	}
}