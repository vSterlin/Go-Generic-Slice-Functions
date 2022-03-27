package gsf

import "testing"

func TestMapIntToInt(t *testing.T) {
	originalSlice := []int{1, 2, 3}

	newSlice := MapSlice(originalSlice, func(x int) int {
		return x * 2
	})

	for i, actual := range newSlice {
		expected := originalSlice[i] * 2
		if actual != expected {
			t.Errorf("expected value to equal: %d, got: %d", expected, actual)
		}
	}
}
