package main

import "fmt"

import "sort"

func binarySearch(nums []int, target int) bool {
	sort.Ints(nums)
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false

}

func main() {
	nums := []int{4, 6, 7, 2, 3, 77, 89, 44, 5, 54, 34, 41, 17, 79, 9, 99, 1}
	target := 18

	fmt.Println("The target at nums is :", binarySearch(nums, target))
}
