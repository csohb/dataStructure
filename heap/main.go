package heap

type Lesser interface {
	Less(v Lesser) bool
}

type Heap struct {
	arr []Lesser
}

func (h *Heap) Len() int {
	return len(h.arr)
}

func (h *Heap) Push(v Lesser) {
	h.arr = append(h.arr, v)

	h.upHeapify(len(h.arr) - 1)
}
func (h *Heap) upHeapify(idx int) {
	parentIdx := (idx - 1) / 2
	if parentIdx < 0 {
		return
	}

	if h.arr[parentIdx].Less(h.arr[idx]) {
		h.arr[parentIdx], h.arr[idx] = h.arr[idx], h.arr[parentIdx]
		h.upHeapify(parentIdx)
	}
}

func (h *Heap) Pop() Lesser {
	root := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.downHeapify(0)

	return root
}

func (h *Heap) downHeapify(idx int) {
	l := len(h.arr)

	leftChildIdx := (idx * 2) + 1
	rightChildIdx := (idx * 2) + 2

	changeIdx := -1
	if leftChildIdx < l && h.arr[idx].Less(h.arr[leftChildIdx]) {
		changeIdx = leftChildIdx
	}

	if rightChildIdx < l && h.arr[idx].Less(h.arr[rightChildIdx]) {
		if changeIdx < 0 || h.arr[changeIdx].Less(h.arr[rightChildIdx]) {
			changeIdx = rightChildIdx
		}
	}

	if changeIdx >= 0 {
		h.arr[idx], h.arr[changeIdx] = h.arr[changeIdx], h.arr[idx]
		h.downHeapify(changeIdx)
	}
}
