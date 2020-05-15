package main

import "fmt"

// 冒泡排序，a表示数组，n表示数组大小
func bubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j] // 交换
				flag = true                 // 表示有数据交换
			}
		}
		if !flag {
			break // 没有数据交换，提前退出
		}
	}
}

func main() {
	nums := []int{2, 4, 77, 3, 99, 23, 4, 1, 8, 42, 14, 56}
	bubbleSort(nums, len(nums))
	fmt.Println(nums)
}
