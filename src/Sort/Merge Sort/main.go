/*
归并排序，小有序逐渐变为大有序，最终解决全部问题
*/

package main

import (
	"fmt"
)

// 自顶向下归并
func mergeSort(nums []int) []int {
	lens := len(nums)
	if lens < 2 {
		return nums
	}
	middle := lens / 2
	left := nums[0:middle]
	right := nums[middle:]
	return merge(mergeSort(left), mergeSort(right))

}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:] //改变切片起始位置，从而不用改变归并的下标，妙啊
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func main() {
	nums := []int{6, 1, 3, 4, 9, 8, 2, 5, 7}

	fmt.Println(mergeSort(nums))
}
