package test_func_type

import (
	"fmt"
	"testing"
)

func TestFuncType(t *testing.T) {
	f := func(i int) {
		fmt.Println("param is", i)
	}
	f(1)
}
