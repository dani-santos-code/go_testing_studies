package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// checkSums := func(t *testing.T, got, want []int) {
	// 	t.Helper()
	// 	if !reflect.DeepEqual(got, want) {
	// 		t.Errorf("got %v want %v", got, want)
	// 	}
	// }

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("sum all numbers given one or more slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("sums tails of a collection", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns 0 if no slices are passed", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d want %q, given %v", got, want, numbers)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
