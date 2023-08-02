package genericgokit

import (
	"testing"
)

func TestFind(t *testing.T) {
	val, ok := Find([]string{"a", "b", "c"}, func(i string) bool {
		arr := []string{"1", "2", "3"}
		for _, v := range arr {
			if v == i {
				return true
			}
		}
		return false
	})
	if ok {
		t.Errorf("Find(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", ok, false)
		return
	}
	if val != "a" {
		t.Errorf("Find(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", val, "a")
		return
	}
}

func TestFindIndexOf(t *testing.T) {
	val, ok, index := FindIndexOf([]string{"a", "b", "c"}, func(i string) bool {
		arr := []string{"1", "2", "3"}
		for _, v := range arr {
			if v == i {
				return true
			}
		}
		return false
	})
	if ok {
		t.Errorf("FindIndexOf(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", ok, false)
		return
	}
	if val != "a" {
		t.Errorf("FindIndexOf(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", val, "a")
		return
	}
	if index != 0 {
		t.Errorf("FindIndexOf(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", index, 0)
		return
	}
}

func TestFindLastIndexOf(t *testing.T) {
	val, ok, index := FindLastIndexOf([]string{"a", "b", "c"}, func(i string) bool {
		arr := []string{"1", "2", "3"}
		for _, v := range arr {
			if v == i {
				return true
			}
		}
		return false
	})
	if ok {
		t.Errorf("FindLastIndexOf(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", ok, false)
		return
	}
	if val != "a" {
		t.Errorf("FindLastIndexOf(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", val, "a")
		return
	}
	if index != 2 {
		t.Errorf("FindLastIndexOf(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", index, 2)
		return
	}
}

func TestFindOrElse(t *testing.T) {
	val := FindOrElse([]string{"a", "b", "c"}, func(i string) bool {
		arr := []string{"1", "2", "3"}
		for _, v := range arr {
			if v == i {
				return true
			}
		}
		return false
	})
	if val != "d" {
		t.Errorf("FindOrElse(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", val, "d")
		return
	}
}
