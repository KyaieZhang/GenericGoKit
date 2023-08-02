package genericgokit

import (
	"math/rand"
	"time"
)

func RandomSort[T any](slice []T) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func BubbleSort[T any](slice []T, less func(a, b T) bool) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if less(slice[j], slice[i]) {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
	}
}

func SelectSort[T any](slice []T, less func(a, b T) bool) {
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if less(slice[j], slice[min]) {
				min = j
			}
			slice[i], slice[min] = slice[min], slice[i]
		}
	}
}

func InsertSort[T any](slice []T, less func(a, b T) bool) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if less(slice[j], slice[j-1]) {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

func QuickSort[T any](slice []T, less func(a, b T) bool) {
	if len(slice) <= 1 {
		return
	}
	left, right := 0, len(slice)-1
	pivot := rand.Int() % len(slice)
	slice[pivot], slice[right] = slice[right], slice[pivot]
	for i := range slice {
		if less(slice[i], slice[right]) {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
		if i == right {
			break
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	QuickSort(slice[:left], less)
	QuickSort(slice[left+1:], less)
}

//func MergeSort[T any](slice []T, less func(a, b T) bool) {
//	if len(slice) <= 1 {
//		return
//	}
//	mid := len(slice) / 2
//	MergeSort(slice[:mid], less)
//	MergeSort(slice[mid:], less)
//	merge(slice, less)
//}
//
//func merge[T any](slice []T, less func(a, b T) bool) {
//	mid := len(slice) / 2
//	left, right := slice[:mid], slice[mid:]
//	a, b := 0, 0
//	for k := 0; k < len(slice); k++ {
//		if a >= len(left) {
//			slice[k] = right[b]
//			b++
//		}
//		if b >= len(right) {
//			slice[k] = left[a]
//			a++
//		}
//	}
//	for i := 0; i < len(slice); i++ {
//		for j := i + 1; j < len(slice); j++ {
//			if less(slice[j], slice[i]) {
//				slice[i], slice[j] = slice[j], slice[i]
//			}
//		}
//	}
//}
//
//func HeapSort[T any](slice []T, less func(a, b T) bool) {
//	if len(slice) <= 1 {
//		return
//	}
//	for i := len(slice) / 2; i >= 0; i-- {
//		heapify(slice, i, len(slice), less)
//	}
//	for i := len(slice) - 1; i > 0; i-- {
//		slice[0], slice[i] = slice[i], slice[0]
//		heapify(slice, 0, i, less)
//	}
//}
//
//func heapify[T any](slice []T, i, n int, less func(a, b T) bool) {
//	largest := i
//	left := 2*i + 1
//	right := 2*i + 2
//	if left < n && less(slice[left], slice[largest]) {
//		largest = left
//		slice[i], slice[largest] = slice[largest], slice[i]
//		heapify(slice, largest, n, less)
//	}
//	if right < n && less(slice[right], slice[largest]) {
//		largest = right
//		slice[i], slice[largest] = slice[largest], slice[i]
//		heapify(slice, largest, n, less)
//	}
//}
