package data_structures

type Node[T any] struct {
	item T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	last *Node[T]
	size int
}

func (list *LinkedList[T]) Add(value T) {
	newNode := Node[T]{
		item: value,
		next: nil,
	}

	if list.head == nil {
		list.head = &newNode
		list.last = &newNode
	} else {
		list.last.next = &newNode
		list.last = &newNode
	}

	list.size++
}

func (list LinkedList[T]) ToArray() []T {
	array := []T{}
	head := list.head

	for head != nil {
		array = append(array, list.head.item)
		head = head.next
	}

	return array
}
