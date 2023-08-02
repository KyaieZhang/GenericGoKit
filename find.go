package genericgokit

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

func FindOrElse[T any](collection []T, predicate func(T) bool) T {
	for _, item := range collection {
		if predicate(item) {
			return item
		}
	}
	var zero T
	return zero
}
