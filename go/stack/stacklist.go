package stack

//StackList implements data structure: stack
type StackList struct {
	top   *item
	count int
}

type item struct {
	next  *item
	value interface{}
}

//New retuns new Stack
func New() *StackList {
	return new(StackList)
}

//Push puts the item on top of the stack
func (s *StackList) Push(v interface{}) {
	s.count++
	newItem := &item{
		value: v,
	}
	newItem.next = s.top
	s.top = newItem
}

//Count returns the count of items in the stack
func (s *StackList) Count() int {
	return s.count
}

//Peek returns a item from the top of the stack, does not extract it
func (s *StackList) Peek() (interface{}, error) {
	if s.top == nil {
		return nil, ErrEmptyStack
	}
	return s.top.value, nil
}

//Pop extracts a item from the top of the stack and returns it
func (s *StackList) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, ErrEmptyStack
	}
	s.count--
	v := s.top
	s.top = v.next
	return v.value, nil
}
