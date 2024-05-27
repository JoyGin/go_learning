package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancel(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

// 【不推荐】发送消息用来关闭 channel，只有一个接受者可以收到
func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

// 【标准操作】通过关闭 channel 实现任务取消
// 关闭 channel 的操作是一个广播通知，所有监听该 channel 的都可以收到通知
func cancel_2(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {
	ch := make(chan struct{})
	for i := 0; i < 4; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancel(ch) {
					break
				} else {
					time.Sleep(time.Millisecond * 5)
				}
			}
			fmt.Println(i, " canceled")
		}(i, ch)
	}
	//cancel_1(ch)
	cancel_2(ch)
	time.Sleep(time.Second * 1)
}
