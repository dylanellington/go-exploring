package data_structures

type CircularQueue[T any] struct {
	data []T
	head int
	tail int
	size int
}

// NewCircularQueue creates and returns a default instance of CircularQueue[T]
func NewCircularQueue[T any](size int) CircularQueue[T] {
	queue := CircularQueue[T]{
		data: make([]T, size),
		head: -1,
		tail: -1,
		size: size,
	}

	return queue
}

// EnQueue adds the given value to the queue.
// Returns false if the queue is full and the value cannot be added.
func (queue *CircularQueue[T]) EnQueue(value T) bool {
	if queue.IsEmpty() {
		queue.head = 0
		queue.tail = 0
		queue.data[queue.tail] = value
		return true
	}

	tailIndex := (queue.tail + 1) % queue.size

	if tailIndex != queue.head {
		queue.tail = tailIndex
		queue.data[queue.tail] = value
		return true
	}

	return false
}

// DeQueue removes and returns the first value from the queue.
// If the queue is empty a default T will be returned.
func (queue *CircularQueue[T]) DeQueue() T {
	var newItem T

	if queue.IsEmpty() {
		return newItem
	}

	headValue := queue.data[queue.head]
	queue.data[queue.head] = newItem

	if queue.head == queue.tail {
		queue.head = -1
		queue.tail = -1
	} else {
		headIndex := (queue.head + 1) % queue.size
		queue.head = headIndex
	}

	return headValue
}

// Peek returns the first item in the queue without removing it.
// If the queue is empty a default T will be returned.
func (queue *CircularQueue[T]) Peek() T {
	var newItem T

	if queue.IsEmpty() {
		return newItem
	}

	return queue.data[queue.head]
}

// PeekRear returns the last item in the queue without removing it.
// If the queue is empty a default T will be returned.
func (queue *CircularQueue[T]) PeekRear() T {
	var newItem T

	if queue.IsEmpty() {
		return newItem
	}

	return queue.data[queue.tail]
}

// IsEmpty returns true if the queue has no items in it.
func (queue *CircularQueue[T]) IsEmpty() bool {
	return queue.head == -1
}

// IsFull returns true if the queue size is at its capacity.
func (queue *CircularQueue[T]) IsFull() bool {
	return queue.head == (queue.tail + 1) % queue.size
}