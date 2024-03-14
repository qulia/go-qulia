package broker

// PubSub messaging broker with generic type
type Broker[T any] interface {
	Subscribe() chan T
	Unsubsribe(ch chan T)
	Publish(msg T)
	Start()
	Stop()
}

func NewBroker[T any](msgBufferSize int) Broker[T] {
	return newBrokerImpl[T](msgBufferSize)
}
