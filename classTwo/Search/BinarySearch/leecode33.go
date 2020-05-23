/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。
*/
package main

import "fmt"

// 先找到最小元素的下标，然后拆成两个有序数组分别查找
func search(nums []int, target int) int {
	s := findSmallest(nums)
	ans1 := bsearch(nums[:s], len(nums[:s]), target)
	ans2 := bsearch(nums[s:], len(nums[s:]), target)
	if ans1 != -1 {
		return ans1
	} else if ans2 != -1 {
		return ans2 + len(nums[:s])
	}
	return -1
}

// 有序数组二分查找
func bsearch(nums []int, n, value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == value {
			return mid
		} else if nums[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 返回最小元素的下标
func findSmallest(nums []int) int {
	l, r := 0, len(nums)-1 //左右下标定位
	for l < r {
		m := (l + r) / 2       //计算中间下标
		if nums[m] > nums[r] { //如果中间值大于右边界，则m之前的元素都不可能在*右排序数组*内，所以缩小左边界
			l = m + 1
		} else if nums[m] < nums[r] { //如果中间值小于右边界，则m之后的元素绝对比*右排序数组*的首个元素大或等于，所以缩小右边界
			r = m
		} else {
			r-- //如果中间值等于右边界，则无法判断m在哪个排序数组中，r缩小一位，从而m缩小一位，在继续判断
		}
	}
	return l
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(nums, 0))
}
