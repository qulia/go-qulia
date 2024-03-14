package broker_test

import (
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/messaging/broker"
	"github.com/stretchr/testify/assert"
)

const (
	MessageBufferSize  = 5
	TestMessageContent = 42
)

const timeout = time.Second

func TestPublish(t *testing.T) {
	broker := broker.NewBroker[int](MessageBufferSize)
	go broker.Start()
	defer broker.Stop()
	subscriber := broker.Subscribe()

	time.Sleep(time.Second) // wait before publishing to ensure subscription is processed
	broker.Publish(TestMessageContent)

	select {
	case msg := <-subscriber: // from buffered
		assert.Equal(t, msg, TestMessageContent)
	case <-time.After(timeout):
		t.Fatal("Timed out waiting for message")
	}
}

func TestBrokerStartStop(t *testing.T) {
	broker := broker.NewBroker[int](MessageBufferSize)
	go broker.Start()

	subscriber := broker.Subscribe()
	broker.Stop()

	_, ok := <-subscriber
	if ok {
		t.Fatal("Expected subscriber channel to be closed")
	}
}

func TestMultiMessageBroadcast(t *testing.T) {
	const (
		numSubscribers = 3
		msgCount       = 10
	)

	msgs := [msgCount]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	broker := broker.NewBroker[int](MessageBufferSize)
	go broker.Start()
	defer broker.Stop()
	donewg := &sync.WaitGroup{}
	subwg := &sync.WaitGroup{}
	for i := 0; i < numSubscribers; i++ {
		donewg.Add(1)
		go func(subwg, donewg *sync.WaitGroup) {
			defer donewg.Done()

			msgCh := broker.Subscribe()
			rcvdMsgs := [msgCount]int{}
			rcvd := 0
			for msg := range msgCh {
				rcvdMsgs[msg-1] = msg
				rcvd++
				if rcvd == msgCount {
					broker.Unsubsribe(msgCh)
				}
			}

			assert.Equal(t, msgs, rcvdMsgs)
		}(subwg, donewg)
	}

	time.Sleep(time.Second)
	// Publishing messages
	for _, val := range msgs {
		broker.Publish(val)
	}

	donewg.Wait()
}
