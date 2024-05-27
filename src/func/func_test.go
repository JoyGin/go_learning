package test_func

import "testing"

// 可变长参数
func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}

func sum(ops ...int) int {
	sum := 0
	for _, op := range ops {
		sum += op
	}
	return sum
}

// defer 函数
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources")
	}()
	t.Log("Started")
	//panic("Fatal error")
}
