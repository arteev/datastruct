package stack

//StackArray implements data structure: stack
type StackArray struct {
	data []interface{}
}

//NewStackArray retuns new StackArray
func NewStackArray() *StackArray {
	return &StackArray{}
}

//Push puts the item on top of the stack
func (s *StackArray) Push(v interface{}) {
	s.data = append(s.data, v)
}

//Count returns the count of items in the stack
func (s *StackArray) Count() int {
	return len(s.data)
}

//Peek returns a item from the top of the stack, does not extract it
func (s *StackArray) Peek() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

//Pop extracts a item from the top of the stack and returns it
func (s *StackArray) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, ErrEmptyStack
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}
