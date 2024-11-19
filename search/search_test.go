package search

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type searcher[T cmp.Ordered] struct {
	name     string
	searchFn func([]T, T) error
}

var searchers = []searcher[int]{
	{
		name:     "LinearSearch",
		searchFn: Linear[int],
	},
	{
		name:     "BinarySearch",
		searchFn: Binary[int],
	},
}

func TestSearch(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected error
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			target:   3,
			expected: nil,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			target:   6,
			expected: ErrNotFound,
		},
		{
			input:    []int{10, 20, 30, 40, 50},
			target:   10,
			expected: nil,
		},
		{
			input:    []int{10, 20, 30, 40, 50},
			target:   50,
			expected: nil,
		},
		{
			input:    []int{},
			target:   1,
			expected: ErrNotFound,
		},
		{
			input:    []int{5},
			target:   5,
			expected: nil,
		},
		{
			input:    []int{5},
			target:   10,
			expected: ErrNotFound,
		},
	}

	for _, searcher := range searchers {
		t.Run(searcher.name, func(t *testing.T) {
			for _, tt := range tests {
				err := searcher.searchFn(tt.input, tt.target)
				assert.Equal(t, tt.expected, err)
			}
		})
	}
}

var benchInput []int

func init() {
	const size = 2000
	benchInput = make([]int, size)

	for i := 0; i < size; i++ {
		benchInput[i] = rand.Intn(1000)
	}

	slices.Sort(benchInput)
}

func BenchmarkSearchers(b *testing.B) {
	for _, searcher := range searchers {
		b.Run(searcher.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = searcher.searchFn(benchInput, benchInput[500])
			}
		})
	}
}
