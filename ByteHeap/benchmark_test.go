// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).
package byteHeap

import (
	"math/rand"
	"testing"

	"github.com/klaytn/klaytn/common"
)

func BenchmarkPrqueByteSlicePush(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([][]byte, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = common.MakeRandomBytes(256)
	}
	// Execute the benchmark
	b.ResetTimer()
	queue := NewByteSlice()
	for i := 0; i < len(data); i++ {
		queue.Push(data[i], prio[i])
	}
}

func BenchmarkPrqueByteSlicePop(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([][]byte, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = common.MakeRandomBytes(256)
	}
	queue := NewByteSlice()
	for i := 0; i < len(data); i++ {
		queue.Push(data[i], prio[i])
	}
	// Execute the benchmark
	b.ResetTimer()
	for !queue.Empty() {
		queue.Pop()
	}
}

func BenchmarkHeapByteSlicePush(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([][]byte, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = common.MakeRandomBytes(256)
	}
	// Execute the benchmark
	b.ResetTimer()
	var queue ByteHeap
	for i := 0; i < len(data); i++ {
		queue.Push(prio[i])
	}
}

func BenchmarkHeapByteSlicePop(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([][]byte, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = common.MakeRandomBytes(256)
	}
	var queue ByteHeap
	for i := 0; i < len(data); i++ {
		queue.Push(prio[i])
	}
	// Execute the benchmark
	b.ResetTimer()
	for queue.Len() > 0 {
		queue.Pop()
	}
}
