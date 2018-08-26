package stack

import "testing"

func TestStackArray(t *testing.T) {
	testStack(NewStackArray(), t)
}

func BenchmarkStackArrayPush(b *testing.B) {
	benchmarkStackPush(NewStackArray(), b)
}

func BenchmarkStackArrayPushPeek(b *testing.B) {
	benchmarkStackPushPeek(NewStackArray(), b)
}

func BenchmarkStackArrayPushPop(b *testing.B) {
	benchmarkStackPushPop(NewStackArray(), b)
}
