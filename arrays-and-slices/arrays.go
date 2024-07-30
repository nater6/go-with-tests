package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAllTails(sumList ...[]int) []int {
	sums := make([]int, 0)
	for _, numbers := range sumList {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		sum := Sum(numbers[1:])
		sums = append(sums, sum)
	}
	return sums
}
