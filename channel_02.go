// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

package main

import (
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		time.Sleep(10 * time.Millisecond) // 模拟一些工作
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		println(num)
	}
}

func main() {
	ch := make(chan int, 10) // 带缓冲的通道，缓冲区大小为10

	go producer(ch)
	go consumer(ch)

	time.Sleep(2 * time.Second) // 等待协程执行完毕
}
