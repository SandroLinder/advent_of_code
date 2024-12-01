package util

import "fmt"

type QueueNode[T any] struct {
	next *QueueNode[T]
	data T
}

func NewQueueNode[T any](data T) *QueueNode[T] {
	return &QueueNode[T]{data: data}
}

type Queue[T any] struct {
	head   *QueueNode[T]
	tail   *QueueNode[T]
	Length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	node := NewQueueNode(value)
	q.Length++
	if q.tail == nil && q.head == nil {
		q.head = node
		q.tail = node
	} else if q.tail != nil {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue[T]) Deque() (*T, error) {
	if q.Length <= 0 {
		return nil, fmt.Errorf("no Elements in Queue. Cannot deque")
	}

	q.Length--

	head := q.head
	q.head = head.next
	return &head.data, nil
}

func (q *Queue[T]) Debug() {
	head := q.head

	for head != nil {
		println(head.data)
		head = head.next
	}
}
