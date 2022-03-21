package data_structures

import (
	"testing"
)

func Test_LinkedList_Add(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	if list.size != 2 {
		t.Errorf("Expected list size of 2 but actual size is %q.", list.size)
	}
}

func Test_LinkedList_ToArray(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	array := list.ToArray()

	if len(array) == 0 {
		t.Errorf("List to array conversion returned an empty array.")
	}

	if array[0] != 1 && array[1] != 2 {
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
