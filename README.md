# golang-heap
Benchmark for Golang heap implementation

# Prque Int

## Run benchmark
```
cd IntHeap
go test -run=None -bench=BenchmarkPop > intprque.txt
go test -run=None -bench=BenchmarkIntHeapPop > intheap.txt
```

## Run with memory usage
```
cd IntHeap
go test -benchmem -run=None -bench=BenchmarkPop_Mem > intprque-mem.txt
go test -benchmem -run=None -bench=BenchmarkIntHeapPop_Mem > intheap-mem.txt
```

# Prque []Byte

## Run benchmark
```
cd ByteHeap
go test -run=None -bench=BenchmarkPrqueByteSlicePop > byteprque.txt
go test -run=None -bench=BenchmarkHeapByteSlicePop > byteheap.txt
```

## Run with memory usage
```
cd ByteHeap
go test -run=None -bench=BenchmarkPrqueByteSlicePop_Mem > byteprque-mem.txt
go test -run=None -bench=BenchmarkHeapByteSlicePop_Mem > byteheap-mem.txt
```