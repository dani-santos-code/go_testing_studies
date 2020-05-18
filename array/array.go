package array

// Sum sums the values inside an array
func Sum(a []int) int {
	var sum int
	for _, number := range a {
		sum += number
	}
	return sum
}
