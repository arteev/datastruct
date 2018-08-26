package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	testStack(New(), t)
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
