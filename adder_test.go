package integers

import "testing"

func TestAdded(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("want %d got %d", expected, sum)
	}
}
