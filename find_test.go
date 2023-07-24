package genericgokit

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIndexOf(t *testing.T) {
	if IndexOf(nil, "a") != -1 {
		t.Errorf("IndexOf(nil, 'a') == %v, expected %v", IndexOf(nil, "a"), -1)
		return
	}
	if IndexOf([]string{}, "a") != -1 {
		t.Errorf("IndexOf([]string{}, 'a') == %v, expected %v", IndexOf([]string{}, "a"), -1)
		return
	}
	if IndexOf([]string{"a"}, "b") != -1 {
		t.Errorf("IndexOf([]string{}, 'a') == %v, expected %v", IndexOf([]string{"a"}, "b"), -1)
		return
	}
	if IndexOf([]string{"a"}, "a") != 0 {
		t.Errorf("IndexOf([]string{}, 'a') == %v, expected %v", IndexOf([]string{"a"}, "a"), 0)
		return
	}
}

func TestLastIndexOf(t *testing.T) {
	if LastIndexOf(nil, "a") != -1 {
		t.Errorf("LastIndexOf(nil, 'a') == %v, expected %v", LastIndexOf(nil, "a"), -1)
		return
	}
	if LastIndexOf([]string{}, "a") != -1 {
		t.Errorf("LastIndexOf([]string{}, 'a') == %v, expected %v", LastIndexOf([]string{}, "a"), -1)
		return
	}
	if LastIndexOf([]string{"a"}, "b") != -1 {
		t.Errorf("LastIndexOf([]string{}, 'a') == %v, expected %v", LastIndexOf([]string{"a"}, "b"), -1)
		return
	}
	if LastIndexOf([]string{"a"}, "a") != 0 {
		t.Errorf("LastIndexOf([]string{}, 'a') == %v, expected %v", LastIndexOf([]string{"a"}, "a"), 0)
		return
	}
	if LastIndexOf([]string{"a", "b", "a"}, "a") != 2 {
		t.Errorf("LastIndexOf([]string{}, 'a') == %v, expected %v", LastIndexOf([]string{"a", "b", "a"}, "a"), 2)
		return
	}
	if LastIndexOf([]string{"a", "b", "a"}, "b") != 1 {
		t.Errorf("LastIndexOf([]string{}, 'a') == %v, expected %v", LastIndexOf([]string{"a", "b", "a"}, "b"), 1)
		return
	}
}

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
	}, "d")
	if val != "d" {
		t.Errorf("FindOrElse(nil, func(i interface{}) bool { return i.(string) == 'a' }) == %v, expected %v", val, "d")
		return
	}
}

func TestFindKey(t *testing.T) {
	val, ok := FindKey(map[string]string{"a": "1", "b": "2", "c": "3"}, "a")
	if ok {

	}
	if val != "a" {
		t.Errorf("FindKey(nil, 'a') == %v, expected %v", val, "a")
		return
	}
}

func TestFindKeyBy(t *testing.T) {
	key, val, ok := FindKeyBy(map[string]string{"a": "1", "b": "2", "c": "3"}, func(a, b string) bool {
		arr := []string{"1", "2", "3"}
		for _, v := range arr {
			if v == a {
				return true
			}
		}
		return false
	})
	if ok {
		t.Errorf("FindKeyBy(nil, func(a, b string) bool { return a == 'a' }) == %v, expected %v", ok, false)
		return
	}
	if key != "a" {
		t.Errorf("FindKeyBy(nil, func(a, b string) bool { return a == 'a' }) == %v, expected %v", key, "a")
		return
	}
	if val != "1" {
		t.Errorf("FindKeyBy(nil, func(a, b string) bool { return a == 'a' }) == %v, expected %v", val, "1")
		return
	}
}

func TestFindUnique(t *testing.T) {
	vals := FindUniques([]string{"a", "b", "c", "a"})
	if len(vals) != 3 {
		t.Errorf("FindUniques(nil, 'a') == %v, expected %v", len(vals), 3)
		return
	}
}

func TestFindUniquesBy(t *testing.T) {
	vals := FindUniquesBy([]string{"a", "b", "c", "a"}, func(a string) bool {
		arr := []string{"1", "2", "3"}
		for _, v := range arr {
			if v == a {
				return true
			}
		}
		return false
	})
	if len(vals) != 3 {
		t.Errorf("FindUniquesBy(nil, func(a string) bool { return a == 'a' }) == %v, expected %v", len(vals), 3)
		return
	}
}

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

func TestGetRandom(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	r := GetRandom([]int{1, 2, 3})
	fmt.Println(fmt.Sprintf("GetRandom([]int{1, 2, 3}) == %v", r))
}

func TestGetRandoms(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	r := GetRandoms([]int{1, 2, 3}, 2)
	fmt.Println(fmt.Sprintf("GetRandoms([]int{1, 2, 3}, 2) == %v", r))
}
