package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("Enqueue", func(t *testing.T) {
		q := New()

		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		assert.Equal(t, 10, q.front.Value)
		assert.Equal(t, 30, q.rear.Value)
		assert.False(t, q.IsEmpty())
		assert.Equal(t, 3, q.Len())
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := New()

		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		elem, err := q.Dequeue()

		assert.NoError(t, err)
		assert.Equal(t, 10, elem)
		assert.Equal(t, 20, q.front.Value)
		assert.Equal(t, 2, q.Len())

		elem, err = q.Dequeue()

		assert.NoError(t, err)
		assert.Equal(t, 20, elem)
		assert.Equal(t, 30, q.front.Value)
		assert.Equal(t, 1, q.Len())

		elem, err = q.Dequeue()

		assert.NoError(t, err)
		assert.Equal(t, 30, elem)
		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Len())

		_, err = q.Dequeue()
		assert.Error(t, err)
	})
}
