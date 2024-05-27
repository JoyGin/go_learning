package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return true
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 4; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancel(ctx) {
					break
				} else {
					time.Sleep(time.Millisecond * 5)
				}
			}
			fmt.Println(i, " canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
