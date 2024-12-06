package util

import (
	"fmt"
)

type StackNode[T comparable] struct {
	prev  *StackNode[T]
	value T
}

func NewStackNode[T comparable](value T) *StackNode[T] {
	return &StackNode[T]{value: value}
}

type Stack[T comparable] struct {
	head   *StackNode[T]
	Length int
}

func NewStack[T comparable]() *Stack[T] {
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

func (s *Stack[T]) UniqueItems() int {
	uniqueItems := make(map[T]bool)

	head := s.head

	for head != nil {
		_, ok := uniqueItems[head.value]

		if !ok {
			uniqueItems[head.value] = true
		}

		head = head.prev
	}

	return len(uniqueItems)
}

func (s *Stack[T]) Contains(elem T) bool {
	head := s.head

	for head != nil {
		if head.value == elem {
			return true
		}

		head = head.prev
	}

	return false
}

func (s *Stack[T]) ToArray() []T {
	head := s.head

	array := make([]T, s.Length)
	index := 0
	for head != nil {
		array[index] = head.value
		head = head.prev
		index++
	}
	return array
}

func (s *Stack[T]) Debug() {
	head := s.head

	for head != nil {
		fmt.Println(head.value)
		head = head.prev
	}
	fmt.Println()
}
