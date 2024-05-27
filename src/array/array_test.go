package array_test

import "testing"

/*
 * 数组声明
 */
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)

	arr5 := [2][2]int{{1, 1}, {2, 2}}
	t.Log(arr5)
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[:3]
	arr3Sec[0] = 0
	t.Log(arr3)
	t.Log(arr3Sec)
}

/*
 * 数组遍历
 */
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}
