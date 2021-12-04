package main

import "testing"

func TestSum(t *testing.T){
	array := [5]int {1,2,3,4,5}
	got := Sum(array) 
	expected := 15
	 
	if got != expected {
		t.Errorf("got %d but expected %d,%v",got,expected,array)
	}
}