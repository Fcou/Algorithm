/*
变体三：查找第一个大于等于给定值的元素
*/
package main

import "fmt"

func bsearch(nums []int, n, value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= value {
			if (mid == 0) || (nums[mid-1] < value) {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 8, 8, 8, 8, 9, 9, 11}
	fmt.Println(bsearch(nums, len(nums), 10))
}
