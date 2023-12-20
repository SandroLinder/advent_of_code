package util

import "fmt"

type Queue struct {
	size     int
	count    int
	elements []interface{}
	front    int
	end      int
}

func NewQueue(size int) *Queue {
	return &Queue{front: 0, end: 0, count: 0, size: size, elements: make([]interface{}, size)}
}

func (q *Queue) Enqueue(element interface{}) error {
	if q.end-q.front == q.size {
		return fmt.Errorf("queue is full")
	}

	q.elements[q.end] = element
	if q.end+1 == q.size {
		q.end = 0
	} else {
		q.end++
	}
	q.count++

	return nil
}

func (q *Queue) Dequeue() interface{} {
	front := q.front
	element := q.elements[front]
	q.elements[front] = nil
	if q.front+1 == q.size {
		q.front = 0
	} else {
		q.front++
	}
	q.count--

	return element
}

func (q *Queue) Print() {
	for i := q.front; i < q.end; i++ {
		fmt.Printf("%v ", q.elements[i])
	}

	fmt.Println()
}

func (q *Queue) CountItemsInQueue() int {
	return q.count
}
