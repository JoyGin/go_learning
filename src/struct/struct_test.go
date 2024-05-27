package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

type Student struct {
	name   string
	gender int
	age    int
}

func (stu *Student) sayHelloPtr(friend string) {
	fmt.Printf("(%x)hello %s, my name is %s\n", unsafe.Pointer(stu), friend, stu.name)
}

func (stu Student) sayHello(friend string) {
	fmt.Printf("(%x)hello %s, my name is %s\n", unsafe.Pointer(&stu), friend, stu.name)
}

func TestStructOperations(t *testing.T) {
	stu := Student{name: "John", gender: 1, age: 18}
	t.Logf("address %x", unsafe.Pointer(&stu))
	stu.sayHello("July")
	stu.sayHelloPtr("July")
}
