// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).

// This is a duplicated and slightly modified version of "gopkg.in/karalabe/cookiejar.v2/collections/prque".

// Package prque implements a priority queue data structure supporting arbitrary
// value types and int64 priorities.
//
// If you would like to use a max-priority queue, use NewByteSlice.
// If you would like to use a min-priority queue, use NewByteSliceInverted.
//
// Internally the queue is based on the standard heap package working on a
// sortable version of the block based stack.

package byteHeap

import "container/heap"

type PrqueByteSlice struct {
	cont *sstackByte
}

// NewByteSlice creates a new max-priority queue.
func NewByteSlice() *PrqueByteSlice {
	return &PrqueByteSlice{newsstackByte(false)}
}

// NewByteSliceInverted creates a new min-priority queue.
func NewByteSliceInverted() *PrqueByteSlice {
	return &PrqueByteSlice{newsstackByte(true)}
}

// Pushes a value with a given priority into the queue, expanding if necessary.
func (p *PrqueByteSlice) Push(data interface{}, priority []byte) {
	heap.Push(p.cont, &itemByteSlice{data, priority})
}

// Pops the value with the greates priority off the stack and returns it.
// Currently no shrinking is done.
func (p *PrqueByteSlice) Pop() (interface{}, []byte) {
	item := heap.Pop(p.cont).(*itemByteSlice)
	return item.value, item.priority
}

// Pops only the item from the queue, dropping the associated priority value.
func (p *PrqueByteSlice) PopItem() interface{} {
	return heap.Pop(p.cont).(*itemByteSlice).value
}

// Checks whether the priority queue is empty.
func (p *PrqueByteSlice) Empty() bool {
	return p.cont.Len() == 0
}

// Returns the number of element in the priority queue.
func (p *PrqueByteSlice) Size() int {
	return p.cont.Len()
}

// Clears the contents of the priority queue.
func (p *PrqueByteSlice) Reset() {
	if p.cont.reverse {
		*p = *NewByteSlice()
	} else {
		*p = *NewByteSliceInverted()
	}
}
