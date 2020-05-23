/*
容易出错的 3 个地方。
1. 循环退出条件注意是 low<=high，而不是 low<high。
2.mid 的取值实际上，mid=(low+high)/2 这种写法是有问题的。因为如果 low 和 high 比较大的话，两者之和就有可能会溢出。
改进的方法是将 mid 的计算方式写成 low+(high-low)/2。更进一步，
如果要将性能优化到极致的话，我们可以将这里的除以 2 操作转化成位运算 low+((high-low)>>1)。因为相比除法运算来说，计算机处理位运算要快得多。
3.low 和 high 的更新low=mid+1，high=mid-1。注意这里的 +1 和 -1，
如果直接写成 low=mid 或者 high=mid，就可能会发生死循环。比如，当 high=3，low=3 时，如果 a[3]不等于 value，就会导致一直循环不退出。

*/
package main

import "fmt"

// 二分查找，基于循环，n表示数字总个数，value表示要查找的数字，找到返回下标，否则返回-1
func bsearchOne(nums []int, n, value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1 // 对于(low + high) / 2的优化
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

// 二分查找的递归实现
func bsearch(nums []int, n, value int) int {
	return bsearchInternally(nums, 0, n-1, value)
}

func bsearchInternally(nums []int, low, high, value int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)>>1
	if nums[mid] == value {
		return mid
	} else if nums[mid] < value {
		return bsearchInternally(nums, mid+1, high, value)
	} else {
		return bsearchInternally(nums, low, mid-1, value)
	}
}

func main() {
	nums := []int{2, 5, 9, 23, 33, 44, 67, 68, 88, 89, 99}
	fmt.Println(bsearch(nums, len(nums), 33))
}
