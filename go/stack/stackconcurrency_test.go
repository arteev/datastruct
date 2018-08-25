package stack

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackConcurrency(t *testing.T) {
	s := NewConcurrency()
	assert.NotNil(t, s)
	_ = Stack(s)
}

func TestConcurency(t *testing.T) {
	stack := NewConcurrency()
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

func BenchmarkStackPushConcurrency(b *testing.B) {
	benchmarkStackPush(NewConcurrency(), b)
}

func BenchmarkStackPushPeekConcurrency(b *testing.B) {
	benchmarkStackPushPeek(NewConcurrency(), b)
}

func BenchmarkStackPushPopConcurrency(b *testing.B) {
	benchmarkStackPushPop(NewConcurrency(), b)
}
