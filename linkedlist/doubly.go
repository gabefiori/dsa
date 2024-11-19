package linkedlist

import (
	"fmt"
	"io"
)

type Doubly[T comparable] struct {
	length int
	Head   *DoublyNode[T]
	Tail   *DoublyNode[T]
}

func NewDoubly[T comparable]() *Doubly[T] {
	return &Doubly[T]{length: 0}
}

func (s *Doubly[T]) Length() int {
	return s.length
}

func (s *Doubly[T]) Search(val T) (*DoublyNode[T], error) {
	if s.Head == nil {
		return nil, ErrNotFound
	}

	curr := s.Head

	for curr != nil {
		if curr.Value == val {
			return curr, nil
		}

		curr = curr.Next
	}

	return nil, ErrNotFound
}

func (s *Doubly[T]) Print(w io.Writer) error {
	if s.Head == nil {
		return nil
	}

	cur := s.Head

	for cur != nil {
		if _, err := fmt.Fprintf(w, "%v\n", cur.Value); err != nil {
			return err
		}

		cur = cur.Next
	}

	return nil
}

func (s *Doubly[T]) Prepend(val T) {
	newNode := NewDoublyNode(val)

	s.length++

	if s.Head == nil {
		s.Head = newNode
		s.Tail = s.Head

		return
	}

	newNode.Next = s.Head
	s.Head.Prev = newNode
	s.Head = newNode
}

func (s *Doubly[T]) Append(val T) {
	newNode := NewDoublyNode(val)

	s.length++

	if s.Head == nil {
		s.Head = newNode
		s.Tail = s.Head

		return
	}

	newNode.Prev = s.Tail
	s.Tail.Next = newNode
	s.Tail = newNode
}

func (s *Doubly[T]) InsertAt(val T, pos int) error {
	if pos > s.length || pos < 0 {
		return ErrInvalidPos
	}

	if pos == s.length {
		s.Append(val)
		return nil
	}

	if pos == 0 {
		s.Prepend(val)
		return nil
	}

	curr := s.Head

	for i := 0; curr != nil && i < pos; i++ {
		curr = curr.Next
	}

	newNode := NewDoublyNode(val)

	newNode.Prev = curr.Prev
	newNode.Next = curr

	curr.Prev.Next = newNode

	s.length++

	return nil
}

func (s *Doubly[T]) DeleteFirst() error {
	if s.Head == nil || s.length == 0 {
		return ErrEmptyList
	}

	if s.length == 1 {
		s.Head = nil
		s.Tail = nil

		s.length--

		return nil
	}

	s.Head.Next.Prev = nil
	s.Head = s.Head.Next

	s.length--

	return nil
}

func (s *Doubly[T]) DeleteLast() error {
	if s.Head == nil || s.length == 0 {
		return ErrEmptyList
	}

	if s.length == 1 {
		s.Head = nil
		s.Tail = nil

		s.length--

		return nil
	}

	s.Tail = s.Tail.Prev
	s.Tail.Next = nil

	s.length--

	return nil
}

func (s *Doubly[T]) DeleteAt(pos int) error {
	if s.Head == nil || s.length == 0 {
		return ErrEmptyList
	}

	if pos > s.length || pos < 0 {
		return ErrInvalidPos
	}

	if pos == s.length-1 {
		return s.DeleteLast()
	}

	if pos == 0 {
		return s.DeleteFirst()
	}

	curr := s.Head

	for i := 0; curr != nil && i < pos; i++ {
		curr = curr.Next
	}

	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev

	s.length--

	return nil
}
