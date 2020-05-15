package main

import "fmt"

// 插入排序，a表示数组，n表示数组大小
func insertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	// 从第二个数开始向前插入
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break // a[j] <= value  之前的数都小
			}
		}
		a[j+1] = value // 插入
	}
}

func main() {
	nums := []int{2, 4, 77, 3, 99, 23, 4, 1, 8, 42, 14, 56}
	insertionSort(nums, len(nums))
	fmt.Println(nums)
}
