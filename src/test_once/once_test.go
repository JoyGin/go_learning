package test_once

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance Singleton
var once sync.Once

func GetSingletonObj() Singleton {
	once.Do(func() {
		fmt.Println("create object")
		singleInstance = Singleton{}
	})
	return singleInstance
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(&obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
