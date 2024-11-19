package stack

import "errors"

var ErrEmptyStack = errors.New("Empty stack")

type Stack[T any] struct {
	elems []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		elems: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.elems = append(s.elems, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroVal T

		return zeroVal, ErrEmptyStack
	}

	idx := len(s.elems) - 1
	elem := s.elems[idx]

	s.elems = s.elems[:idx]

	return elem, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroVal T

		return zeroVal, ErrEmptyStack
	}

	return s.elems[len(s.elems)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.elems)
}
