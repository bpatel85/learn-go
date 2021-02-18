package common

type Queue interface {
	Enqueue(val interface{})
	Dequeue() (interface{}, error)
	IsEmpty() bool
	HasNext() bool
}
