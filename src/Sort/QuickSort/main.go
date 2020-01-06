/*
append 是个坑
算法思想很简单，但数组下标判断很烦，改天再想
*/
package main

import (
	"fmt"
)

func QuickSort(nums []int) []int {
	lens := len(nums)
	var less []int
	var big []int
	var landp []int
	if lens < 2 {
		return nums
	} else {
		pivot := 0
		small := 0

		for i := 1; i < lens; i++ {

			if nums[i] <= pivot {
				small++
				nums[small], nums[i] = nums[i], nums[small]

			}

		}
		if small == 0 {
			big = QuickSort(nums[1:])
			landp = Combinepivot(less, nums[pivot])

			return Combine(landp, big)
		} else if small == lens-1 {
			less = QuickSort(nums[1 : small+1])
			big = QuickSort(nums[small+1:])
			landp = Combinepivot(less, nums[pivot])

			return Combine(landp, big)
		}
	}
}

func Combine(a []int, b []int) []int {
	lena := len(a)
	lenb := len(b)
	s := make([]int, lena+lenb)
	for index := 0; index < lena; index++ {
		s[index] = a[index]
	}
	for index := 0; index < lenb; index++ {
		s[lena+index] = b[index]
	}
	return s
}
func Combinepivot(a []int, b int) []int {
	lena := len(a)
	s := make([]int, lena+1)
	for index := 0; index < lena; index++ {
		s[index] = a[index]
	}

	s[lena] = b
	return s
}
func main() {
	nums := []int{6, 1}

	QuickSort(nums)

	fmt.Println(nums)
}
