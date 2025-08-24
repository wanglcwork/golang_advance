// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func main() {
	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go increment(&wg)
	}

	wg.Wait()
	println("Final Counter:", counter)
	time.Sleep(1 * time.Second) // 确保所有输出完成
}
