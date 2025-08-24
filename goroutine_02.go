// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

package main

import (
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	// 模拟任务执行时间
	time.Sleep(time.Duration(100+id*10) * time.Millisecond)
	elapsed := time.Since(start)
	println("Task", id, "completed in", elapsed.Milliseconds(), "ms")
}

func main() {
	var wg sync.WaitGroup
	numTasks := 10

	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	println("All tasks completed")
}
