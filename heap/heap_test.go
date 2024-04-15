package heap

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type MyInt int

func (m MyInt) Less(v Lesser) bool {
	o, ok := v.(MyInt)
	if !ok {
		panic(" v should be MyInt type")
	}
	return m < o
}

func TestHeapPush(t *testing.T) {
	h := &Heap{}
	for i := 1; i <= 10; i++ {
		h.Push(MyInt(i))
	}

	assert.Equal(t, MyInt(10), h.arr[0].(MyInt))

	max := h.Pop()
	assert.Equal(t, MyInt(10), max)
	assert.Equal(t, MyInt(9), h.arr[0].(MyInt))
}

func TestHeapSort(t *testing.T) {
	h := &Heap{}
	for i := 0; i <= 1000; i++ {
		h.Push(MyInt(rand.Intn(10000)))
	}

	m := h.Pop()
	for len(h.arr) > 0 {
		v := h.Pop()
		assert.False(t, m.Less(v))
	}
}

type MinInt int

func (m MinInt) Less(v Lesser) bool {
	o, ok := v.(MinInt)
	if !ok {
		panic(" v should be MyInt type")
	}
	return m > o
}

func TestHeapTop5(t *testing.T) {
	h := &Heap{}
	for i := 1; i <= 10; i++ {
		h.Push(MinInt(i))
		if h.Len() > 5 {
			h.Pop()
		}
	}

	assert.Equal(t, MinInt(6), h.arr[0].(MinInt))
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

}

func TestHeapTop5_WithGoHeap(t *testing.T) {
	h := &IntHeap{}
	for i := 1; i <= 10; i++ {
		heap.Push(h, i)
		if h.Len() > 5 {
			heap.Pop(h)
		}
	}
	assert.Equal(t, 6, (*h)[0])
}
