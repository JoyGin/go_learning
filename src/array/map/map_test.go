package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m1 := map[string]int{}
	m1["one"] = 1
	m2 := make(map[string]int, 10)

	t.Log(m)
	t.Log(m1)
	t.Log(m2)
}

/**
 * map key 不存在会返回默认值
 */
func TestAccessNotExistKey(t *testing.T) {
	m := map[string]int{}
	if v, ok := m["one"]; !ok {
		t.Log("the key doesn't value, default value", v)
	}
	m["one"] = 1
	if v, ok := m["one"]; ok {
		t.Log("the key has value", v)
	}
}

// map 遍历
func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		t.Log(k, "=", v)
	}
}

// map value 函数使用
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(int) int{}
	m[1] = func(i int) int {
		return i
	}
	m[2] = func(i int) int {
		return i * i
	}
	m[3] = func(i int) int {
		return i * i * i
	}

	t.Log(m[1](2), m[2](2), m[3](2))
}

// map 实现 set
func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	n := 1
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	set[n] = true

	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
