package collection

func Sum(numbers [5]int) int {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	return sum
}
