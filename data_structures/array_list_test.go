package data_structures

import (
	"testing"
)

func Test_ArrayList_Add(t *testing.T) {
	list := ArrayList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	if len(list.array) != 3 {
		t.Errorf("Expected list size of 3 but actual size is %q.", len(list.array))
	}
}

func Test_ArrayList_Remove_First(t *testing.T) {
	list := ArrayList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	list.RemoveAt(0)

	if len(list.array) != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", len(list.array))
	}

	if list.array[0] != itemTwo && list.array[1] != itemThree {
		t.Errorf("Removal caused items in the array to be out of expected order.")
	}
}

func Test_ArrayList_Remove_Middle(t *testing.T) {
	list := ArrayList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	list.RemoveAt(1)

	if len(list.array) != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", len(list.array))
	}

	if list.array[0] != itemOne && list.array[1] != itemThree {
		t.Errorf("Removal caused items in the array to be out of expected order.")
	}
}

func Test_ArrayList_Remove_Last(t *testing.T) {
	list := ArrayList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	list.RemoveAt(2)

	if len(list.array) != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", len(list.array))
	}

	if list.array[0] != itemOne && list.array[1] != itemTwo {
		t.Errorf("Removal caused items in the array to be out of expected order.")
	}
}

func Test_ArrayList_Remove_OutOfBounds(t *testing.T) {
	list := ArrayList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	didRemove := list.RemoveAt(4)

	if didRemove {
		t.Errorf("Remove function returned true when the index to remove was out of bounds.")
	}
}

func Test_ArrayList_Performance(t *testing.T) {
	list := ArrayList[int]{}
	iterations := 100000

	for count := 0; count < iterations; count++ {
		list.Add(count)
	}

	for count := iterations - 1; count >= 0; count-- {
		list.RemoveAt(0)
	}

	if len(list.array) != 0 {
		t.Errorf("Expected list size of 0 but actual size is %q.", len(list.array))
	}
}