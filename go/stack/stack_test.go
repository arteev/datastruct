package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New()
	_ = Stack(s)
	assert.NotNil(t, s)

	v1 := "value1"
	s.Push(v1)
	assert.Equal(t, 1, s.Count())

	v2 := 2
	s.Push(v2)
	assert.Equal(t, 2, s.Count())

	v3 := struct{}{}
	s.Push(v3)

	top, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, v3, top)

	topAgain, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, top, topAgain)

	pop1, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, v3, pop1)
	assert.Equal(t, 2, s.Count())

	pop2, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, v2, pop2)
	assert.Equal(t, 1, s.Count())

	pop3, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, v1, pop3)
	assert.Equal(t, 0, s.Count())

	pop, err := s.Peek()
	assert.Nil(t, pop)
	assert.EqualError(t, err, ErrEmptyStack.Error())

	pop, err = s.Pop()
	assert.Nil(t, pop)
	assert.EqualError(t, err, ErrEmptyStack.Error())

}

var result interface{}

func benchmarkStackPush(s Stack, b *testing.B) {
	v1 := "{}"
	for i := 0; i < b.N; i++ {
		s.Push(v1)
	}
}
func benchmarkStackPushPeek(s Stack, b *testing.B) {

	v1 := "{}"
	for i := 0; i < b.N; i++ {
		s.Push(v1)
		r, _ := s.Peek()
		result = r
	}
}

func benchmarkStackPushPop(s Stack, b *testing.B) {
	v1 := "{}"
	for i := 0; i < b.N; i++ {
		s.Push(v1)
		r, _ := s.Pop()
		result = r
	}
}

func BenchmarkStackPush(b *testing.B) {
	benchmarkStackPush(New(), b)
}

func BenchmarkStackPushPeek(b *testing.B) {
	benchmarkStackPushPeek(New(), b)
}

func BenchmarkStackPushPop(b *testing.B) {
	benchmarkStackPushPop(New(), b)
}
