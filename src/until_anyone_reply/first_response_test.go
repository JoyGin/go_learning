package until_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", i)
}

func FirstResponse() string {
	numOfRunner := 10
	//ch := make(chan string)
	// 预防携程泄漏
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	fmt.Printf("before goroutine: %d\n", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Printf("after goroutine: %d\n", runtime.NumGoroutine())
}
