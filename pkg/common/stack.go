package common

type Stack interface {
	Push(val interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	HasNext() bool
}
