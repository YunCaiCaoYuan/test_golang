package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	cancelCtx, _ := context.WithCancel(ctx)
	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	//cancelFunc()
	time.Sleep(time.Second * 1)
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		// Done 方法返回一个 Channel 对象。在 Context 被取消时，此 Channel 会被 close，如果没被取消，可能会返回 nil。
		// 关于 Done 方法 ：如果 Done 没有被 close，Err 方法返回 nil；如果 Done 被 close，Err 方法会返回 Done 被 close 的原因。
		// 思考：如果 Done 如果没有被 close，会一直阻塞在这里
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
