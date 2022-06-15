package data_structures

import "testing"

func Test_CircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue[int](2)
	q.EnQueue(1)

	if q.Peek() != 1 {
		t.Errorf("EnQueue failed to add item to the queue.")
	}
}

func Test_CircularQueue_DeQueue(t *testing.T) {
	q := NewCircularQueue[int](2)
	q.EnQueue(5)
	value := q.DeQueue()

	if value != 5 {
		t.Errorf("DeQueue failed to get an item from the queue.")
	}
}

func Test_CircularQueue_PeekRear(t *testing.T) {
	q := NewCircularQueue[int](2)
	q.EnQueue(1)
	q.EnQueue(2)

	if q.PeekRear() != 2 {
		t.Errorf("Queue peek rear failed.")
	}
}

func Test_CircularQueue_IsEmpty(t *testing.T) {
	q := NewCircularQueue[int](1)

	if !q.IsEmpty() {
		t.Errorf("Queue is empty check failed.")
	}
}

func Test_CircularQueue_IsFull(t *testing.T) {
	q := NewCircularQueue[int](2)
	q.EnQueue(1)
	q.EnQueue(2)

	if !q.IsFull() {
		t.Errorf("Queue is full check failed.")
	}
}
