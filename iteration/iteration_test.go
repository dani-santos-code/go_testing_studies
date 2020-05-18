package iteration

import (
	"fmt"
	"testing"
)

func TestIteratingCharacter(t *testing.T) {

	t.Run("repeats a single character N times", func(t *testing.T) {
		got := RepeatCharacter("a", 10)
		want := "aaaaaaaaaa"
		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}

	})
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatCharacter("a", 5)
	}
}

func ExampleRepeatCharacter() {
	r := RepeatCharacter("H", 5)
	fmt.Printf("%q\n", r)
	// Output: "HHHHH"
}
