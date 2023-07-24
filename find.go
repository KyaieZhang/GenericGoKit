package genericgokit

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func IndexOf[T comparable](collection []T, item T) int {
	for i, v := range collection {
		if v == item {
			return i
		}
	}
	return -1
}

func LastIndexOf[T comparable](collection []T, item T) int {
	for i := len(collection) - 1; i >= 0; i-- {
		if collection[i] == item {
			return i
		}
	}
	return -1
}

func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func FindIndexOf[T any](collection []T, predicate func(T) bool) (T, bool, int) {
	for i, item := range collection {
		if predicate(item) {
			return item, true, i
		}
	}
	var zero T
	return zero, false, -1
}

func FindLastIndexOf[T any](collection []T, predicate func(T) bool) (T, bool, int) {
	for i := len(collection) - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], true, i
		}
	}
	var zero T
	return zero, false, -1
}

func FindOrElse[T any](collection []T, predicate func(T) bool, orElse T) T {
	for _, item := range collection {
		if predicate(item) {
			return item
		}
	}
	var zero T
	return zero
}

func FindKey[K comparable, V any](collection map[K]V, key K) (V, bool) {
	value, ok := collection[key]
	return value, ok
}

func FindKeyBy[K comparable, V any](collection map[K]V, predicate func(K, V) bool) (K, V, bool) {
	for key, value := range collection {
		if predicate(key, value) {
			return key, value, true
		}
	}
	var zero K
	var zeroV V
	return zero, zeroV, false
}

func FindUniques[T comparable](collection []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, item := range collection {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
			continue
		}
	}
	return result
}

func FindUniquesBy[T any, U comparable](collection []T, predicate func(item T) U) []T {
	seen := make(map[U]bool)
	result := make([]T, 0)
	for _, item := range collection {
		key := predicate(item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
			continue
		}
	}
	return result
}

func Min[T constraints.Ordered](collection []T) T {
	var min T
	if len(collection) == 0 {
		return min
	}
	min = collection[0]
	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item < min {
			min = item
		}
	}
	return min
}

func Max[T constraints.Ordered](collection []T) T {
	var max T
	if len(collection) == 0 {
		return max
	}
	max = collection[0]
	for i := 1; i < len(collection); i++ {
		item := collection[i]
		if item > max {
			max = item
		}
	}
	return max
}

func GetRandom[T any](collection []T) T {
	if len(collection) == 0 {
		var zero T
		return zero
	}
	return collection[rand.Intn(len(collection))]
}

func GetRandoms[T any](collection []T, count int) []T {
	size := len(collection)

	ts := append([]T{}, collection...)

	var results []T

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i

		index := rand.Intn(size - i)
		results = append(results, ts[index])

		ts[index] = ts[copyLength-1]
		ts = ts[:copyLength-1]
	}

	return results
}
