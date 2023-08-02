package genericgokit

import (
	"testing"
)

func TestMin(t *testing.T) {
	if Min([]int{1, 2, 3}) != 1 {
		t.Errorf("Min([]int{1, 2, 3}) == %v, expected %v", Min([]int{1, 2, 3}), 1)
		return
	}
}

func TestMax(t *testing.T) {
	if Max([]int{1, 2, 3}) != 3 {
		t.Errorf("Max([]int{1, 2, 3}) == %v, expected %v", Max([]int{1, 2, 3}), 3)
		return
	}
}
