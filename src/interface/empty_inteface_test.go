package interface_test

import (
	"fmt"
	"testing"
)

// go 空接口类似 c 中的 void* 类型，可以指向任何类型
func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown type", v)
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
