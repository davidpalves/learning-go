package arrayandslices

// Sum receives a list of integers and return its elements sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
