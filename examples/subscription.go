package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	. "pubsub-test/util"
	"time"
)

func main() {
	var (
		ctx    = context.Background()
		client = MustGetPubSubClient(ctx)
		sub    = client.Subscription("trades-sub")
		delay  = time.Second
	)
	defer client.Close()

	log.Printf("receiving from test-sub")
	ReceiveMessage(ctx, sub, func(_ context.Context, msg *pubsub.Message) {
		var (
			msgContent = string(msg.Data)
			msgAttempt = 0
		)
		if msg.DeliveryAttempt != nil {
			msgAttempt = *msg.DeliveryAttempt
		}
		LogMessage(msgContent, fmt.Sprintf("attempt %d", msgAttempt))
		LogMessage(msgContent, "processing")
		time.Sleep(delay)
		LogMessage(msgContent, "processed")
		if msgContent == "test 7" && msgAttempt < 10 {
			msg.Nack()
			LogMessage(msgContent, "return nack")
			return
		}
		msg.Ack()
		LogMessage(msgContent, "return ack")
	})
}
