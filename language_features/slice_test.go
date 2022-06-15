package language_features

import (
	"testing"
)

func Test_Slice_SharedMemory(t *testing.T) {
	x := make([]int, 0, 4) // array (x), capacity of 4
	x = append(x, 1, 2, 3) // add 3 numbers to (x), 3/4 memory spaces used
	// x = [1, 2, 3]

	y := x[:1] // array (y), references first memory space in (x)
	y = append(y, 4) // modify (y), changes referenced memory space in (x) at index 2
	// x = [1, 4, 3], y = [1, 4]

	yExpanded := y[:3] // expand (y) memory to the size of (x)
	// yExpanded = [1, 4, 3]

	for index := range yExpanded {
		if yExpanded[index] != x[index] {
			t.Errorf("Shared memory slices x and y (expanded) are not equal: %v, %v\n", x, yExpanded)
		}
	}

	y = append(y, 5, 6, 7) // append to (y) past its capacity to force an internal copy of the array to new memory
	x[0], x[1], x[2] = 10, 11, 12 // changes to (x) wont effect (y) now
	// x = 10, 11, 12, y = 1, 4, 5, 6, 7

	if x[0] == y[0] || x[1] == y[1] || x[2] == y[2] {
		t.Errorf("Non shared memory slices x and y have equal values: %v, %v\n", x, y)
	}
}
