package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
	. "pubsub-test/util"
)

func main() {
	var (
		ctx    = context.Background()
		client = MustGetPubSubClient(ctx)
		sub    = client.Subscription("traders_error-sub")
	)
	defer client.Close()

	log.Printf("receiving from traders_error-sub")
	ReceiveMessage(ctx, sub, func(_ context.Context, msg *pubsub.Message) {
		msgContent := string(msg.Data)
		LogMessage(msgContent, "received")
		msg.Ack()
	})
}
