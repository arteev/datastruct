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

//StackBase implements data structure: stack
type StackBase struct {
	top   *item
	count int
}

type item struct {
	next  *item
	value interface{}
}

//New retuns new Stack
func New() *StackBase {
	return new(StackBase)
}

//Push puts the item on top of the stack
func (s *StackBase) Push(v interface{}) {
	s.count++
	newItem := &item{
		value: v,
	}
	newItem.next = s.top
	s.top = newItem
}

//Count returns the count of items in the stack
func (s *StackBase) Count() int {
	return s.count
}

//Peek returns a item from the top of the stack, does not extract it
func (s *StackBase) Peek() (interface{}, error) {
	if s.top == nil {
		return nil, ErrEmptyStack
	}
	return s.top.value, nil
}

//Pop extracts a item from the top of the stack and returns it
func (s *StackBase) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, ErrEmptyStack
	}
	s.count--
	v := s.top
	s.top = v.next
	return v.value, nil
}
