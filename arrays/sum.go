package main

// Adds all elements of a slice and returns the total
func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

// Takes many slices, finds the sum of all the elements, and returns the sums of each slice as a single slice
func SumAll(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

// A tail is a collection of all items except for the first one i.e head. This function takes the numbers of the tail and sums them
func SumAllTails(slices ...[]int) []int {
	var sumTails []int

	for _, slice := range slices {
		tail := slice[1:]
		sumTails = append(sumTails, Sum(tail))
	}

	return sumTails
}
