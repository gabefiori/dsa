package search

import "cmp"

func Linear[T cmp.Ordered](arr []T, target T) error {
	for _, v := range arr {
		if v == target {
			return nil
		}
	}

	return ErrNotFound
}
