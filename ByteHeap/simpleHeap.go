// This shows an example
package byteHeap

import "bytes"

type ByteHeap [][]byte

func (h ByteHeap) Len() int {
	return len(h)
}

func (h ByteHeap) Less(i, j int) bool {
	return bytes.Compare(h[i], h[j]) > 0
}

func (h ByteHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ByteHeap) Push(element interface{}) {
	*h = append(*h, element.([]byte))
}

func (h *ByteHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}
