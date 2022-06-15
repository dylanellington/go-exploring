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

// NewLinkedList creates and returns a default instance of LinkedList[T]
func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T] {
		head: nil,
		last: nil,
		size: 0,
	}
}

// Add inserts a given item at the end of the list.
func (list *LinkedList[T]) Add(value T) {
	newNode := Node[T]{
		item: value,
		next: nil,
	}

	if list.head == nil {
		list.head = &newNode
		list.last = list.head
	} else {
		list.last.next = &newNode
		list.last = list.last.next
	}

	list.size++
}

// RemoveAt removes an item from the list at a given index.
func (list *LinkedList[T]) RemoveAt(index int) bool {
	if index < 0 || index >= list.size {
		return false
	}

	if index == 0 {
		list.head = list.head.next
		list.size--
		return true
	}

	node := list.head
	nodeToRemove := node.next

	for i := 1; i < index; i++ {
		node = nodeToRemove
		nodeToRemove = nodeToRemove.next
	}

	node.next = nodeToRemove.next
	list.size--

	return true
}

// ToArray converts the list to an array.
func (list LinkedList[T]) ToArray() []T {
	array := []T{}
	head := list.head

	for head != nil {
		array = append(array, list.head.item)
		head = head.next
	}

	return array
}
