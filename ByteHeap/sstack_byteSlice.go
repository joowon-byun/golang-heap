// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).

// This is a duplicated and slightly modified version of "gopkg.in/karalabe/cookiejar.v2/collections/prque".

package byteHeap

import "bytes"

const blockSize = 4096

// A prioritized item in the sorted stack.
//
// Note: priorities can "wrap around" the int64 range, a comes before b if bytes.Compare(a.priority - b.priority) > 0.
// The difference between the lowest and highest priorities in the queue at any point should be less than 2^63.
type itemByteSlice struct {
	value    interface{}
	priority []byte
}

// Internal sortable stack data structure. Implements the Push and Pop ops for
// the stack (heap) functionality and the Len, Less and Swap methods for the
// sortability requirements of the heaps.
type sstackByte struct {
	size     int
	capacity int
	offset   int
	reverse  bool // reverse the result of Less()

	blocks [][]*itemByteSlice
	active []*itemByteSlice
}

// Creates a new, empty stack.
func newsstackByte(reverse bool) *sstackByte {
	result := new(sstackByte)
	result.active = make([]*itemByteSlice, blockSize)
	result.blocks = [][]*itemByteSlice{result.active}
	result.capacity = blockSize
	result.reverse = reverse
	return result
}

// Pushes a value onto the stack, expanding it if necessary. Required by
// heap.Interface.
func (s *sstackByte) Push(data interface{}) {
	if s.size == s.capacity {
		s.active = make([]*itemByteSlice, blockSize)
		s.blocks = append(s.blocks, s.active)
		s.capacity += blockSize
		s.offset = 0
	} else if s.offset == blockSize {
		s.active = s.blocks[s.size/blockSize]
		s.offset = 0
	}
	s.active[s.offset] = data.(*itemByteSlice)
	s.offset++
	s.size++
}

// Pops a value off the stack and returns it. Currently no shrinking is done.
// Required by heap.Interface.
func (s *sstackByte) Pop() (res interface{}) {
	s.size--
	s.offset--
	if s.offset < 0 {
		s.offset = blockSize - 1
		s.active = s.blocks[s.size/blockSize]
	}
	res, s.active[s.offset] = s.active[s.offset], nil
	return
}

// Returns the length of the stack. Required by sort.Interface.
func (s *sstackByte) Len() int {
	return s.size
}

// Compares the priority of two elements of the stack (higher is first).
// Required by sort.Interface.
func (s *sstackByte) Less(i, j int) bool {
	r := bytes.Compare(s.blocks[i/blockSize][i%blockSize].priority, s.blocks[j/blockSize][j%blockSize].priority) > 0
	if s.reverse {
		return !r
	}
	return r
}

// Swaps two elements in the stack. Required by sort.Interface.
func (s *sstackByte) Swap(i, j int) {
	ib, io, jb, jo := i/blockSize, i%blockSize, j/blockSize, j%blockSize
	s.blocks[ib][io], s.blocks[jb][jo] = s.blocks[jb][jo], s.blocks[ib][io]
}

// Resets the stack, effectively clearing its contents.
func (s *sstackByte) Reset() {
	*s = *newsstackByte(s.reverse)
}
