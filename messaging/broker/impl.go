package broker

type brokerImpl[T any] struct {
	subcribers    map[chan T]bool
	stopCh        chan bool
	subCh         chan chan T
	unsubCh       chan chan T
	publishCh     chan T
	msgBufferSize int
}

func (b *brokerImpl[T]) Publish(msg T) {
	b.publishCh <- msg
}

func (b *brokerImpl[T]) Subscribe() chan T {
	msgCh := make(chan T, b.msgBufferSize)
	b.subCh <- msgCh
	return msgCh
}

func (b *brokerImpl[T]) Unsubsribe(ch chan T) {
	b.unsubCh <- ch
}

func (b *brokerImpl[T]) Start() {
	for {
		select {
		case <-b.stopCh:
			for msgCh := range b.subcribers {
				close(msgCh)
			}
			b.subcribers = make(map[chan T]bool)
			return
		case msgCh := <-b.subCh:
			b.subcribers[msgCh] = true
		case msgCh := <-b.unsubCh:
			delete(b.subcribers, msgCh)
			close(msgCh)
		case msg := <-b.publishCh:
			for rcvCh := range b.subcribers {
				// non-blocking send
				go func(msgCh chan T) {
					rcvCh <- msg
				}(rcvCh)
			}
		}
	}
}

func (b *brokerImpl[T]) Stop() {
	b.stopCh <- true
}

func newBrokerImpl[T any](msgBufferSize int) Broker[T] {
	return &brokerImpl[T]{
		subcribers:    make(map[chan T]bool),
		stopCh:        make(chan bool),   // unbuffered
		subCh:         make(chan chan T), // unbuffered
		unsubCh:       make(chan chan T), // unbuffered
		publishCh:     make(chan T, 1),
		msgBufferSize: msgBufferSize,
	}
}
