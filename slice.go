package genericgokit

func SliceContains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func SliceContainsWithFunc[T any](slice []T, f func(T) bool) bool {
	for _, s := range slice {
		if f(s) {
			return true
		}
	}
	return false
}

func SliceRemoveElement[T comparable](slice []T, item T) []T {
	writeIdx := 0
	for _, val := range slice {
		if val != item {
			slice[writeIdx] = val
			writeIdx++
		}
	}
	return slice[:writeIdx]
}

func SliceRemoveElementWithFunc[T any](slice []T, f func(T) bool) []T {
	writeIdx := 0
	for _, val := range slice {
		if !f(val) {
			slice[writeIdx] = val
			writeIdx++
			continue
		}
	}
	return slice[:writeIdx]
}

func SliceUnique[T comparable](slice []T) []T {
	writeIdx := 0
	for _, val := range slice {
		if !SliceContains(slice[:writeIdx], val) {
			slice[writeIdx] = val
			writeIdx++
			continue
		}
	}
	slice = slice[:writeIdx]
	return slice
}

func SliceUniqueWithFunc[T any](slice []T, f func(T) bool) []T {
	writeIdx := 0
	for _, val := range slice {
		if !SliceContainsWithFunc(slice[:writeIdx], f) {
			slice[writeIdx] = val
			writeIdx++
			continue
		}
	}
	slice = slice[:writeIdx]
	return slice
}

func SliceReplaceElement[T comparable](slice []T, old T, new T) []T {
	for i, s := range slice {
		if s == old {
			slice[i] = new
		}
	}
	return slice
}

func SliceReplaceElementWithFunc[T any](slice []T, f func(T) bool, new T) []T {
	for i, s := range slice {
		if f(s) {
			slice[i] = new
			continue
		}
	}
	return slice
}
