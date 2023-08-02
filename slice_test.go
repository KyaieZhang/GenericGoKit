package genericgokit

import (
	"testing"
)

func TestSliceContains(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if !SliceContains(slice, "a") {
		t.Error("Slice should contain 'a'")
	}
}

func TestSliceContainsWithFunc(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if !SliceContainsWithFunc(slice, func(s string) bool {
		return s == "a"
	}) {
		t.Error("Slice should contain 'a'")
	}
}

func TestSliceRemoveElement(t *testing.T) {
	slice := []string{"a", "b", "c"}
	SliceRemoveElement(slice, "a")
	if SliceContains(slice, "a") {
		t.Error("Slice should not contain 'a'")
	}
}

func TestSliceRemoveElementWithFunc(t *testing.T) {
	slice := []string{"a", "b", "c"}
	SliceRemoveElementWithFunc(slice, func(s string) bool {
		return s == "a"
	})
	if SliceContains(slice, "a") {
		t.Error("Slice should not contain 'a'")
	}
}

func TestSliceUnique(t *testing.T) {
	slice := []string{"a", "b", "c", "a", "b", "c"}
	SliceUnique(slice)
	if len(slice) != 3 {
		t.Error("Slice should have 3 elements")
	}
}

func TestSliceUniqueWithFunc(t *testing.T) {
	slice := []string{"a", "b", "c", "a", "b", "c"}
	SliceUniqueWithFunc(slice, func(s string) bool {
		return s == "a"
	})
	if len(slice) != 2 {
		t.Error("Slice should have 2 elements")
	}
}

func TestSliceReplaceElement(t *testing.T) {
	slice := []string{"a", "b", "c"}
	SliceReplaceElement(slice, "a", "d")
	if !SliceContains(slice, "d") {
		t.Error("Slice should contain 'd'")
	}
}

func TestSliceReplaceElementWithFunc(t *testing.T) {
	slice := []string{"a", "b", "c"}
	SliceReplaceElementWithFunc(slice, func(s string) bool {
		return s == "a"
	}, "a")
	if !SliceContains(slice, "a") {
		t.Error("Slice should contain 'a'")
	}
}
