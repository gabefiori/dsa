package bstree

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		bst := New[int]()

		bst.Insert(100)
		bst.Insert(50)
		bst.Insert(150)
		bst.Insert(125)

		assert.Equal(t, 100, bst.Root.Value)
		assert.Equal(t, 50, bst.Root.Left.Value)
		assert.Equal(t, 150, bst.Root.Right.Value)
		assert.Equal(t, 125, bst.Root.Right.Left.Value)
	})

	t.Run("Search", func(t *testing.T) {
		bst := New[int]()

		bst.Insert(100)
		bst.Insert(50)
		bst.Insert(150)
		bst.Insert(125)

		n := bst.Search(125)

		assert.NotNil(t, n)
		assert.Equal(t, n.Value, 125)
	})

	t.Run("Delete", func(t *testing.T) {
		bst := New[int]()

		bst.Insert(100)

		bst.Insert(50)
		bst.Insert(60)
		bst.Insert(40)

		bst.Insert(150)
		bst.Insert(125)
		bst.Insert(160)
		bst.Insert(155)

		bst.Delete(150)
		bst.Delete(50)

		assert.Equal(t, 155, bst.Root.Right.Value)
		assert.Equal(t, 60, bst.Root.Left.Value)

		bst.Delete(100)
		assert.Equal(t, 125, bst.Root.Value)
	})

	t.Run("InOrder", func(t *testing.T) {
		bst := New[int]()

		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(4)
		bst.Insert(6)
		bst.Insert(7)

		var buf bytes.Buffer
		err := bst.InOrder(&buf)

		assert.NoError(t, err)

		expected := "3\n4\n5\n6\n7\n"
		assert.Equal(t, expected, buf.String())
	})
}
