package main

func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int {
	// numsToSum length tells me how many slices we have
	argLength := len(numsToSum)

	// I need to make an empty slice for the return so that it's accessible outside of the loop
	returnSlice := make([]int, argLength)

	for i, slice := range numsToSum {
		// I can call the function I already made to sum the slice
		returnSlice[i] = Sum(slice)
	}

	return returnSlice
}
