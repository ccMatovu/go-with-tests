package main

import "fmt"

func Sum(numbers []int) int{
	var sum int;
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(slicesToSum ...[]int) []int{
	fmt.Println("checkckckckck")
	slice := []int{}

	for _,v := range slicesToSum{
		var sum int
		for _, value := range v {
			sum += value
		}

		slice = append(slice,sum )

	}
	return slice
}