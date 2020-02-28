/*
append 是个坑，切片扩容丢失信息
算法思想很简单，但数组下标判断很烦，改天再想
解决切片下标办法，之间弃用append,自己写个需要的数组拼接函数
以下实现比之前利用left,right下标来回移动，好理解，同时逻辑判断简单，但多用了内存空间
*/
package main

import (
	"fmt"
)

func QuickSort(nums []int) []int {
	lens := len(nums)
	if lens < 2 {
		return nums
	}
	var less []int
	var big []int

	pivot := nums[0]

	for i := 1; i < lens; i++ {
		if nums[i] <= pivot {
			less = append(less, nums[i])
		} else {
			big = append(big, nums[i])
		}
	}

	return comberArray(comberInt(QuickSort(less), pivot), QuickSort(big))
}

func comberInt(a []int, b int) []int {

	c := make([]int, len(a)+1)
	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	c[len(a)] = b
	return c
}

func comberArray(a []int, b []int) []int {

	c := make([]int, len(a)+len(b))
	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	for j := 0; j < len(b); j++ {
		c[len(a)+j] = b[j]
	}
	return c
}

func main() {
	nums := []int{6, 1, 3, 4, 9, 8, 2, 5, 7, 99, 47, 87, 34}

	fmt.Println(QuickSort(nums))
}
