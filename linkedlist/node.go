package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{Value: val}
}

type DoublyNode[T comparable] struct {
	Value T
	Next  *DoublyNode[T]
	Prev  *DoublyNode[T]
}

func NewDoublyNode[T comparable](val T) *DoublyNode[T] {
	return &DoublyNode[T]{Value: val}
}
