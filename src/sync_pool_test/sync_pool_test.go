package sync_pool_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// sync.Pool 适合用于通过复用，降低复杂对象的创建和 GC 代价
// 协程安全，会有锁的开销
// 生命周期受 GC 影响，不适合做连接池等，需自己管理生命周期的资源的池化

var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("new object")
		return 100
	},
}

func TestSyncPool(t *testing.T) {
	//pool.Put(10)
	//pool.Put(10)
	//pool.Put(10)

	pool.Get()
	pool.Get()
	pool.Get()
	runtime.GC()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(pool.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}
