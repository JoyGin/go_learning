package slice_test

import "testing"

/**
 * 切片声明
 */
func TestSliceInit(t *testing.T) {
	var s0 []int
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	//s := []int{}
	//t.Log(len(s), cap(s))

	s1 := []int{1, 2, 3}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 2, 4)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
