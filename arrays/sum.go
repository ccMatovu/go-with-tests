package main


func Sum(numbers []int) int{
	var sum int;
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(slicesToSum ...[]int) []int{
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

func SumAllTails(slicesToSum ...[]int)[]int{
	slice := []int{}

	for _,v := range slicesToSum{
		
		var sum int
		for  index, value := range v {
			if index == 0{
			continue
		}
			sum += value
		}
		slice = append(slice,sum )
	}
	return slice
}