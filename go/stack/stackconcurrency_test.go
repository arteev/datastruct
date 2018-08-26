package stack

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackConcurrency(t *testing.T) {
	s := NewConcurrency(New())
	assert.NotNil(t, s)
	_ = Stack(s)
}

func TestConcurencyStackOnList(t *testing.T) {
	stack := NewConcurrency(New())
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stack.Push("val")
			_, err := stack.Peek()
			assert.NoError(t, err)
			_, err = stack.Pop()
			assert.NoError(t, err)
		}()
	}
	wg.Wait()

	assert.Equal(t, 0, stack.Count())
}

func BenchmarkStackOnListPushConcurrency(b *testing.B) {
	benchmarkStackPush(NewConcurrency(New()), b)
}

func BenchmarkStackOnListPushPeekConcurrency(b *testing.B) {
	benchmarkStackPushPeek(NewConcurrency(New()), b)
}

func BenchmarkStackOnListPushPopConcurrency(b *testing.B) {
	benchmarkStackPushPop(NewConcurrency(New()), b)
}

func BenchmarkStackOnArrayPushConcurrency(b *testing.B) {
	benchmarkStackPush(NewConcurrency(NewStackArray()), b)
}

func BenchmarkStackOnArrayPushPeekConcurrency(b *testing.B) {
	benchmarkStackPushPeek(NewConcurrency(NewStackArray()), b)
}

func BenchmarkStackOnArrayPushPopConcurrency(b *testing.B) {
	benchmarkStackPushPop(NewConcurrency(NewStackArray()), b)
}
