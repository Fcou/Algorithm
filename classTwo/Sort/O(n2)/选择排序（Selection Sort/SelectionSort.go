package main

import "fmt"

// 选择排序，a表示数组，n表示数组大小
func selectionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	// 每次选择一个最小的，排在前面
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j //记录最小下标
			}
		}
		a[i], a[min] = a[min], a[i] //交换
	}
}

func main() {
	nums := []int{2, 4, 77, 3, 99, 23, 4, 1, 8, 42, 14, 56}
	selectionSort(nums, len(nums))
	fmt.Println(nums)
}
