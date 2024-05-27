package test_select

import (
	"fmt"
	"testing"
	"time"
)

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		fmt.Println("do async service")
		//time.Sleep(500 * time.Millisecond)
		time.Sleep(50 * time.Millisecond)
		retCh <- "hello world"
		fmt.Println("end do async service")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(100 * time.Millisecond):
		t.Error("time out")
	}
}
