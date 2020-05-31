/*
变体一：查找第一个值等于给定值的元素
*/
package main

import "fmt"

func bsearch(nums []int, n, value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == value {
			if (nums[mid-1] != value) || (mid == 0) { // 如果mid之前已经不存在重复元素，则确定返回，否则，进一步缩小范围
				return mid
			} else {
				high = mid - 1
			}
		} else if nums[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 8, 8, 8, 8, 9, 9, 11}
	fmt.Println(bsearch(nums, len(nums), 8))
}
