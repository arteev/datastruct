package stack

import "sync"

//StackConcurrency implements the concurrency safe stack
type StackConcurrency struct {
	mu sync.RWMutex
	Stack
}

//NewConcurrency returns new StackConcurrency
func NewConcurrency() *StackConcurrency {
	return &StackConcurrency{
		Stack: New(),
	}
}

//Push puts the item on top of the stack
func (s *StackConcurrency) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Stack.Push(v)
}

//Count returns the count of items in the stack
func (s *StackConcurrency) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Stack.Count()
}

//Peek returns a item from the top of the stack, does not extract it
func (s *StackConcurrency) Peek() (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Stack.Peek()
}

//Pop extracts a item from the top of the stack and returns it
func (s *StackConcurrency) Pop() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Stack.Pop()
}
