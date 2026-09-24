package main

import (
	"context"
	"log"
	"time"

	"food_delivery/pubsub"
	"food_delivery/pubsub/local_pubsub"
)

func main() {
	var localPS pubsub.Pubsub = local_pubsub.NewLocalPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub_1, _ := localPS.Subscribe(context.Background(), topic)
	sub_2, _ := localPS.Subscribe(context.Background(), topic)

	localPS.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPS.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Con1:", (<-sub_1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Con1:", (<-sub_2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	//close_1()
	//
	//localPS.Publish(context.Background(), topic, pubsub.NewMessage(3))
	//
	//time.Sleep(time.Second * 2)
}
