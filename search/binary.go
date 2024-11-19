package search

import "cmp"

func Binary[T cmp.Ordered](arr []T, target T) error {
	low, high := 0, len(arr)

	for low < high {
		mid := low + (high-low)/2
		v := arr[mid]

		if v == target {
			return nil
		}

		if target > v {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return ErrNotFound
}
