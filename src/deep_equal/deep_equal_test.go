package deep_equal

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	strmap1 := map[int]string{1: "one", 2: "two", 3: "three"}
	strmap2 := map[int]string{1: "one", 2: "two", 3: "three"}
	strmap3 := map[int]string{1: "one", 2: "two", 4: "three"}

	//t.Log(strmap1 == strmap2)
	t.Log(reflect.DeepEqual(strmap1, strmap2))
	t.Log(reflect.DeepEqual(strmap1, strmap3))

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 3, 4}
	//t.Log(slice1 == slice2)
	t.Log(reflect.DeepEqual(slice1, slice2))
	t.Log(reflect.DeepEqual(slice1, slice3))
}
