package data_structures

type ArrayList[T any] struct {
	array []T
}

// NewArrayList creates and returns a default instance of ArrayList[T]
func NewArrayList[T any]() ArrayList[T] {
	return ArrayList[T]{
		array: make([]T, 0 , 10),
	}
}

// Add appends a given item into the rear of the list.
func (list *ArrayList[T]) Add(item T) {
	list.array = append(list.array, item)
}

// RemoveAt removes an item at the given index from the list.
// Returns true if an item was removed.
func (list *ArrayList[T]) RemoveAt(index int) bool {
	if index < 0 || index > len(list.array) - 1 {
		return false
	}

	list.array = append(list.array[0:index], list.array[index + 1:len(list.array)]...)
	return true
}
