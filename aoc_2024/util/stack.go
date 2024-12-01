package util

import "fmt"

type StackNode[T any] struct {
	prev  *StackNode[T]
	value T
}

func NewStackNode[T any](value T) *StackNode[T] {
	return &StackNode[T]{value: value}
}

type Stack[T any] struct {
	head   *StackNode[T]
	Length int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Length: 0,
	}
}

func (s *Stack[T]) Push(value T) {
	node := NewStackNode(value)
	s.Length++

	if s.head == nil {
		s.head = node
	} else {
		node.prev = s.head
		s.head = node
	}
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.head == nil {
		return nil, fmt.Errorf("no element on stack. Cannot pop")
	}

	s.Length--

	node := s.head
	s.head = node.prev

	return &node.value, nil
}

func (s *Stack[T]) Debug() {
	head := s.head

	for head != nil {
		println(head.value)
		head = head.prev
	}
}
