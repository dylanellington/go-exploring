package data_structures

import (
	"testing"
)

func Test_LinkedList_Add(t *testing.T) {
	list := LinkedList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	if list.size != 3 {
		t.Errorf("Expected list size of 3 but actual size is %q.", list.size)
	}
}

func Test_LinkedList_Remove_First(t *testing.T) {
	list := LinkedList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	list.RemoveAt(0)
	array := list.ToArray()

	if list.size != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", list.size)
	}

	if array[0] != itemTwo && array[1] != itemThree {
		t.Errorf("Removal caused items in the array to be out of expected order.")
	}
}

func Test_LinkedList_Remove_Middle(t *testing.T) {
	list := LinkedList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	list.RemoveAt(1)
	array := list.ToArray()

	if list.size != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", list.size)
	}

	if array[0] != itemOne && array[1] != itemThree {
		t.Errorf("Removal caused items in the array to be out of expected order.")
	}
}

func Test_LinkedList_Remove_Last(t *testing.T) {
	list := LinkedList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	list.RemoveAt(2)
	array := list.ToArray()

	if list.size != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", list.size)
	}

	if array[0] != itemOne && array[1] != itemTwo {
		t.Errorf("Removal caused items in the array to be out of expected order.")
	}
}

func Test_LinkedList_Remove_OutOfBounds(t *testing.T) {
	list := LinkedList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	didRemove := list.RemoveAt(4)

	if didRemove {
		t.Errorf("Remove function returned true when the index to remove was out of bounds.")
	}
}

func Test_LinkedList_ToArray(t *testing.T) {
	list := LinkedList[int]{}
	itemOne, itemTwo, itemThree := 1, 2, 3
	list.Add(itemOne)
	list.Add(itemTwo)
	list.Add(itemThree)

	array := list.ToArray()

	if len(array) == 0 {
		t.Errorf("List to array conversion returned an empty array.")
	}

	if array[0] != itemOne && array[1] != itemTwo && array[2] != itemThree {
		t.Errorf("List to array conversion did not contain the expected values.")
	}
}

func Test_LinkedList_ToArray_Empty(t *testing.T) {
	list := LinkedList[int]{}
	array := list.ToArray()

	if len(array) != 0 {
		t.Errorf("List to array conversion did not return an empty array.")
	}
}

func Test_LinkedList_Performance(t *testing.T) {
	list := LinkedList[int]{}
	iterations := 1000

	for count := 0; count < iterations; count++ {
		list.Add(count)
	}

	for count := iterations - 1; count >= 0; count-- {
		list.RemoveAt(count)
	}

	if list.size != 0 {
		t.Errorf("Expected list size of 0 but actual size is %q.", list.size)
	}
}