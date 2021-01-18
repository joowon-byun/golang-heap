// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).
package intHeap

import (
	"math/rand"
	"testing"
)

func BenchmarkPush(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([]int64, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = rand.Int63()
	}
	// Execute the benchmark
	b.ResetTimer()
	queue := New()
	for i := 0; i < len(data); i++ {
		queue.Push(data[i], prio[i])
	}
}

func BenchmarkPop(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([]int64, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = rand.Int63()
	}
	queue := New()
	for i := 0; i < len(data); i++ {
		queue.Push(data[i], prio[i])
	}
	// Execute the benchmark
	b.ResetTimer()
	for !queue.Empty() {
		queue.Pop()
	}
}

func BenchmarkIntHeapPush(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([]int64, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = rand.Int63()
	}
	// Execute the benchmark
	b.ResetTimer()
	var queue IntHeap
	for i := 0; i < len(data); i++ {
		queue.Push(prio[i])
	}
}

func BenchmarkIntHeapPop(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	prio := make([]int64, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
		prio[i] = rand.Int63()
	}
	var queue IntHeap
	for i := 0; i < len(data); i++ {
		queue.Push(prio[i])
	}
	// Execute the benchmark
	b.ResetTimer()
	for queue.Len() > 0 {
		queue.Pop()
	}
}
