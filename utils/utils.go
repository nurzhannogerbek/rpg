package utils

// Sum calculates the sum of integers in an array.
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
