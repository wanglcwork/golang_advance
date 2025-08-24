// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

package main

import (
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) // 模拟一些工作
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		println(num)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(2 * time.Second) // 等待协程执行完毕
}
