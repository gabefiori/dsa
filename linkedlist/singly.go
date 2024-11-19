package linkedlist

import (
	"fmt"
	"io"
)

type Singly[T comparable] struct {
	length int
	Head   *Node[T]
}

func NewSingly[T comparable]() *Singly[T] {
	return &Singly[T]{length: 0}
}

func (s *Singly[T]) Length() int {
	return s.length
}

func (s *Singly[T]) Search(val T) (*Node[T], error) {
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

func (s *Singly[T]) Print(w io.Writer) error {
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

func (s *Singly[T]) Prepend(val T) {
	newNode := NewNode(val)

	s.length++

	if s.Head == nil {
		s.Head = newNode
		return
	}

	newNode.Next = s.Head
	s.Head = newNode
}

func (s *Singly[T]) Append(val T) {
	newNode := NewNode(val)

	s.length++

	if s.Head == nil {
		s.Head = newNode

		return
	}

	curr := s.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func (s *Singly[T]) InsertAt(val T, pos int) error {
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

	for i := 0; curr != nil && i < pos-1; i++ {
		curr = curr.Next
	}

	newNode := NewNode(val)

	newNode.Next = curr.Next
	curr.Next = newNode

	s.length++

	return nil
}

func (s *Singly[T]) DeleteFirst() error {
	if s.Head == nil || s.length == 0 {
		return ErrEmptyList
	}

	if s.length == 1 {
		s.Head = nil

		s.length--

		return nil
	}

	s.Head = s.Head.Next
	s.length--

	return nil
}

func (s *Singly[T]) DeleteLast() error {
	if s.Head == nil || s.length == 0 {
		return ErrEmptyList
	}

	if s.length == 1 {
		s.Head = nil
		s.length--

		return nil
	}

	curr := s.Head
	prev := curr

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil

	s.length--

	return nil
}

func (s *Singly[T]) DeleteAt(pos int) error {
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

	for i := 0; curr != nil && i < pos-1; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	s.length--

	return nil
}
