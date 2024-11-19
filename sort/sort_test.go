package sort

import (
	"cmp"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sorter[T cmp.Ordered] struct {
	name   string
	sortFn func([]T)
}

var sorters = []sorter[int]{
	{
		name:   "InsertionSort",
		sortFn: InsertionSort[int],
	},
	{
		name:   "BubbleSort",
		sortFn: BubbleSort[int],
	},
	{
		name:   "QuickSort",
		sortFn: QuickSort[int],
	},
}

func TestSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{20, 15, 10, 5, 0},
			expected: []int{0, 5, 10, 15, 20},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{-1, -2, -3, -4, -5},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			input:    []int{0},
			expected: []int{0},
		},
		{
			input:    []int{5, 3, 8, 1, 4},
			expected: []int{1, 3, 4, 5, 8},
		},
		{
			input:    []int{100, 50, 0, -50, -100},
			expected: []int{-100, -50, 0, 50, 100},
		},
	}

	for _, sorter := range sorters {
		t.Run(sorter.name, func(t *testing.T) {
			for _, tt := range tests {
				data := make([]int, len(tt.input))
				copy(data, tt.input)

				sorter.sortFn(data)
				assert.Equal(t, tt.expected, data)
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
}

func BenchmarkSorts(b *testing.B) {
	for _, sorter := range sorters {
		b.Run(sorter.name, func(b *testing.B) {
			input := make([]int, len(benchInput))

			for i := 0; i < b.N; i++ {
				copy(input, benchInput)
				sorter.sortFn(input)
			}
		})
	}
}
