package array

// Sum sums the values inside an array
func Sum(a []int) int {
	var sum int
	for _, number := range a {
		sum += number
	}
	return sum
}

//SumAll will sum all values inside an array
func SumAll(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		if Sum(numbers) != 0 {
			sums = append(sums, Sum(numbers))
		}
	}

	return sums
}

// SumAllTails will sum all values inside an array in the tails
func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
