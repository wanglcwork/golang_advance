// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

package main

import (
	"time"
)

func printOdd() {
	for i := 1; i <= 10; i += 2 {
		println(i)
		time.Sleep(100 * time.Millisecond) // 模拟一些工作
	}
}

func printEven() {
	for i := 2; i <= 10; i += 2 {
		println(i)
		time.Sleep(100 * time.Millisecond) // 模拟一些工作
	}
}

func main() {
	go printOdd()
	go printEven()

	time.Sleep(2 * time.Second) // 等待协程执行完毕
}
