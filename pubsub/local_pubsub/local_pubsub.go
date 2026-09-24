package local_pubsub

import (
	"context"
	"log"
	"sync"

	"food_delivery/common"

	"food_delivery/pubsub"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (pb *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, message *pubsub.Message) error {
	message.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		pb.messageQueue <- message
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (<-chan *pubsub.Message, func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	log.Print("ok")

	return c, func() {
		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	go func() {
		defer common.AppRecover()
		for {
			mess := <-ps.messageQueue

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
