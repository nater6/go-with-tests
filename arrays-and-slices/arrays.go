package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}