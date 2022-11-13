package sort

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CallWithTimeout(fn func(), timeout time.Duration) bool {
	ch := make(chan bool)
	defer close(ch)

	go func() {
		fn()
		ch <- true
	}()

	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

func CallWithTimeout2(fn func(), timeout time.Duration) bool {
	ch := make(chan bool)
	defer close(ch)

	go func() {
		fn()
		ch <- true
	}()

	go func() {
		time.Sleep(timeout)
		ch <- false
	}()

	return <-ch
}

func TestBubbleSort(t *testing.T) {
	got := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !CallWithTimeout2(func() { BubbleSort(got) }, 100*time.Millisecond) {
		assert.Fail(t, "Error. The test was running too long")
	}

	assert.Equal(t, want, got)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
		BubbleSort(got)
	}
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
		StdSort(got)
	}
}
