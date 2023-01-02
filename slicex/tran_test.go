package slicex

import "testing"

func TestArrayToMap(t *testing.T) {
	arr := []int{1, 4, 67, 8658, 68685, 7547}
	m := ArrayToMap(arr)
	t.Logf("%+v", m)
}

func TestMapToArray(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 5
	m["c"] = 6
	m["d"] = 9

	arr := MapToArray(m)
	t.Logf("%+v", arr)
}
