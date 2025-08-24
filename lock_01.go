// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

package main

import (
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go increment(&wg, &mu)
	}

	wg.Wait()
	println("Final Counter:", counter)
	time.Sleep(1 * time.Second) // 确保所有输出完成
}
