package genericgokit

import (
	"fmt"
	"testing"
)

func TestRandomSort(t *testing.T) {
	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Printf("old : %v\n", a)
	RandomSort(a)
	fmt.Printf("new : %v\n", a)
}

func TestBubbleSort(t *testing.T) {
	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Printf("old : %v\n", a)
	BubbleSort(a, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("new : %v\n", a)
}

func TestSelectSort(t *testing.T) {
	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Printf("old : %v\n", a)
	SelectSort(a, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("new : %v\n", a)
}

func TestInsertSort(t *testing.T) {
	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Printf("old : %v\n", a)
	InsertSort(a, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("new : %v\n", a)
}

func TestQuickSort(t *testing.T) {
	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Printf("old : %v\n", a)
	QuickSort(a, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("new : %v\n", a)
}

//func TestMergeSort(t *testing.T) {
//	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
//	fmt.Printf("old : %v\n", a)
//	MergeSort(a, func(i, j int) bool {
//		return i > j
//	})
//	fmt.Printf("new : %v\n", a)
//}
//
//func TestHeapSort(t *testing.T) {
//	var a = []int{5, 3, 8, 4, 2, 7, 1, 6}
//	fmt.Printf("old : %v\n", a)
//	HeapSort(a, func(i, j int) bool {
//		return i > j
//	})
//	fmt.Printf("new : %v\n", a)
//}
