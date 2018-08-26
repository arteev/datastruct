package stack

import (
	"errors"
)

// Implements stack using list

//Errors
var (
	ErrEmptyStack = errors.New("stack empty")
)

//Stack data structure
type Stack interface {
	Push(v interface{})
	Count() int
	Peek() (interface{}, error)
	Pop() (interface{}, error)
}

