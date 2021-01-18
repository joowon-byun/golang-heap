# golang-heap
Benchmark for Golang heap implementation

# Result
일반 benchmark
* 그냥 benchmark를 돌릴 때는 그냥 `[]int`보다 `[][]int`가 더 빠르다.
  (`[][]int`은 할당하다가 죽어서 속도 차이 확인 불가하지만, 이미 총 걸린 시간에 차이가 많이난다.)
* `[]int` 을 사용하면 `Out of memory`가 발생한다.

일반 benchmark 결과 분석
* slice가 다 차면 `len(slice) * 2`의 slice를 새로 할당해 주고, 내용물을 옮겨준다.
이 때, `[]int`에 굉장히 많은 값을 넣게 되면, 굉장히 많은 memory가 사용될 수 있다.
* go benchmark는 돌리는 컴퓨터의 cpu, memory에 따라 실행 횟수를 정한다.
(예상) `[]int` 타입이라고 생각하고 실행 횟수를 정했는데, 예상한 것보다 많은 메모리를 사용하고 있어서 out of memory가 발생한게 아닐까?

메모리 사용량과 함께 benchmark
* benchmark를 실행할 때, cpu (`-cpuprofile`) 또는 메모리 사용량(`-benchmem`, `-memprofile`)을 함께 측정할 수 있다.
* 메모리 사용량을 측정하면 `[]int`와 `[][]int`의 사용량이 같으며, `[]int`가 `[][]int`보다 빠르다.

메모리 사용량 benchmark 결과 분석
* 이런 경우 go GC가 지속적으로 실행되며, 포인터(slice 포함)가 많을수록 GC에서 많은 자원을 사용한다.
* 메모리 사용량 결과는 GC 직후에 실행되기 때문에, GC 되기 전의 결과는 보이지 않는다. 즉, go에서는 같은 메모리를 사용하고 있다는 결과를 낸다.
* Slice에 할당하는 크기가 커질수록 Go GC에서 사용하는 자원이 많기 때문에, (Go에서 보여주는 결과가 아닌, go test process에서 사용하는) 실제 총 메모리 사용량을 측정하면 `[][]int`가 `[]int`보다 오래 걸리거나 많은 메모리를 사용한다는 결과가 나올 수 있다.

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
