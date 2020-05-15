package main

import "fmt"

// 归并排序，自顶向下归并
func mergeSort(nums []int) []int {
	lens := len(nums)
	if lens < 2 {
		return nums
	}
	middle := lens / 2
	return merge(mergeSort(nums[0:middle]), mergeSort(nums[middle:]))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:] //改变切片起始位置，从而不用改变归并的下标
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		result = append(result, left...)
	}

	if len(right) != 0 {
		result = append(result, right...)
	}

	return result
}

func main() {
	nums := []int{2, 4, 77, 3, 99, 23, 4, 1, 8, 42, 14, 56}

	fmt.Println(mergeSort(nums))
}
