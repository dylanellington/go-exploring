package data_structures

type CircularQueue struct {
	data []int
	head int
	tail int
	size int
}

func NewCircularQueue(size int) CircularQueue {
	queue := CircularQueue{
		data: make([]int, size),
		head: -1,
		tail: -1,
		size: size,
	}

	return queue
}

func (queue *CircularQueue) EnQueue(value int) bool {
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

func (queue *CircularQueue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}

	queue.data[queue.head] = 0

	if queue.head == queue.tail {
		queue.head = -1
		queue.tail = -1
	} else {
		headIndex := (queue.head + 1) % queue.size
		queue.head = headIndex
	}

	return true
}

func (queue *CircularQueue) Front() int {
	if queue.IsEmpty() {
		return -1
	}

	return queue.data[queue.head]
}

func (queue *CircularQueue) Rear() int {
	if queue.IsEmpty() {
		return -1
	}

	return queue.data[queue.tail]
}

func (queue *CircularQueue) IsEmpty() bool {
	return queue.head == -1
}

func (queue *CircularQueue) IsFull() bool {
	return queue.head == (queue.tail + 1) % queue.size
}