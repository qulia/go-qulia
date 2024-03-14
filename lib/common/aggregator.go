package common

type Aggregator[T any, R any] interface {
	Add(T)
	Remove(T)
	Result() R
}
