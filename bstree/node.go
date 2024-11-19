package bstree

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T cmp.Ordered](val T) *Node[T] {
	return &Node[T]{Value: val}
}
