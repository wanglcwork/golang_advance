// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

package main

func multiplyByTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	multiplyByTwo(&nums)
	for _, num := range nums {
		println(num)
	}
}
