package test_channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 向关闭的 channel 写入数据会导致 panic
// 从关闭的 channel 读取数据，返回的 value 会为默认零值，而 ok 返回 false

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	//var wg sync.WaitGroup
	dataProducer(ch, &wg)
	dataReceiver(ch, &wg)
	dataReceiver(ch, &wg)
	wg.Wait()
}
