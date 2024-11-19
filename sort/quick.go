package sort

import "cmp"

func QuickSort[T cmp.Ordered](arr []T) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort[T cmp.Ordered](arr []T, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
}

func partition[T cmp.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	i++
	arr[high], arr[i] = arr[i], arr[high]

	return i
}
