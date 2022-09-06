package main

import (
	"context"
	"fmt"
	. "pubsub-test/util"
	"sync"
)

func main() {
	var (
		ctx    = context.Background()
		client = MustGetPubSubClient(ctx)
		topic  = client.Topic("trades")
		msgQty = 10
		wg     sync.WaitGroup
	)
	defer client.Close()

	wg.Add(msgQty)
	for i := 0; i < msgQty; i++ {
		i := i
		go func() {
			PublishMessage(ctx, topic, fmt.Sprintf("test %d", i))
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("done")
}
