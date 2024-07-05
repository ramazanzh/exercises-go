package problem11

import (
	"sort"
)

func keysAndValues[T int | string, V any](input map[T]V) ([]T, []V) {
	keys := make([]T, 0, len(input))
	values := make([]V, 0, len(input))

	for key, _ := range input {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, value := range keys {
		values = append(values, input[value])
	}

	return keys, values
}
