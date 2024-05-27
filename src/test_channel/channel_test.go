package test_channel

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	fmt.Println("start do service")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("end do service")
	return "service done"
}

func otherTask() {
	fmt.Println("start do other task")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("end do other task")
}

func TestService(t *testing.T) {
	ret := service()
	fmt.Println(ret)
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

func MyAsyncService() chan string {
	retCh := make(chan string)
	go func() {
		fmt.Println("do my async service")
		time.Sleep(100 * time.Millisecond)
		retCh <- "hello world"
		fmt.Println("end do my async service")
	}()
	return retCh
}

func MyTask() {
	fmt.Println("do my task")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("end do my task")
}

func TestMyAsyncService(t *testing.T) {
	retCh := MyAsyncService()
	MyTask()
	fmt.Println(<-retCh)
}
