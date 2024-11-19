package bstree

import (
	"cmp"
	"fmt"
	"io"
)

type Tree[T cmp.Ordered] struct {
	Root *Node[T]
}

func New[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Insert(val T) {
	t.Root = Insert(t.Root, val)
}

func (t *Tree[T]) Delete(val T) {
	t.Root = Delete(t.Root, val)
}

func (t *Tree[T]) Search(val T) *Node[T] {
	return Search(t.Root, val)
}

func (t *Tree[T]) InOrder(w io.Writer) error {
	return InOrder(t.Root, w)
}

func Insert[T cmp.Ordered](root *Node[T], val T) *Node[T] {
	if root == nil {
		return NewNode(val)
	}

	if root.Value == val {
		return root
	}

	if val > root.Value {
		root.Right = Insert(root.Right, val)
	} else {
		root.Left = Insert(root.Left, val)
	}

	return root
}

func Delete[T cmp.Ordered](root *Node[T], val T) *Node[T] {
	if root == nil {
		return root
	}

	if val > root.Value {
		root.Right = Delete(root.Right, val)
	} else if val < root.Value {
		root.Left = Delete(root.Left, val)
	} else {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		succ := findSuccessor(root)

		root.Value = succ.Value
		root.Right = Delete(root.Right, succ.Value)
	}

	return root
}

func findSuccessor[T cmp.Ordered](curr *Node[T]) *Node[T] {
	curr = curr.Right

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func Search[T cmp.Ordered](root *Node[T], val T) *Node[T] {
	if root == nil || root.Value == val {
		return root
	}

	if val > root.Value {
		return Search(root.Right, val)
	}

	return Search(root.Left, val)
}

func InOrder[T cmp.Ordered](root *Node[T], w io.Writer) error {
	if root == nil {
		return nil
	}

	if err := InOrder(root.Left, w); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "%v\n", root.Value); err != nil {
		return err
	}

	return InOrder(root.Right, w)
}
