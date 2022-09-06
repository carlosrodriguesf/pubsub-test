package util

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
)

func PublishMessage(ctx context.Context, topic *pubsub.Topic, msg string) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("can't publish: %v\n", err)
	}
	log.Printf("published a message: %s (%v)\n", msg, id)
}

func ReceiveMessage(ctx context.Context, sub *pubsub.Subscription, handler func(context.Context, *pubsub.Message)) {
	err := sub.Receive(ctx, handler)
	if err != nil {
		log.Printf("receive error: %v\n", err)
	}
}

func MustGetPubSubClient(ctx context.Context) *pubsub.Client {
	projectID, credentialsOpt := mustGetCredentials()
	client, err := pubsub.NewClient(ctx, projectID, credentialsOpt)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func mustGetCredentials() (string, option.ClientOption) {
	file, err := os.Open("pubsub-credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	credentials, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := new(struct {
		ProjectID string `json:"project_id"`
	})

	err = json.Unmarshal(credentials, config)
	if err != nil {
		log.Fatal(err)
	}

	return config.ProjectID, option.WithCredentialsJSON(credentials)
}
