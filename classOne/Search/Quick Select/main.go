/*
在计算机科学中，快速选择（英語：Quickselect）是一种从无序列表找到第k小元素的选择算法。 它从原理上来说与快速排序有关。 与快速排序一样都由托尼·霍尔提出的，因而也被称为霍尔选择算法。
同样地，它在实际应用是一种高效的算法，具有很好的平均时间复杂度，然而最坏时间复杂度则不理想。
时间复杂度为O(n) ,因为如果我们把每次分区遍历的元素个数加起来，就是：n+n/2+n/4+n/8+…+1。
这是一个等比数列求和，最后的和等于 2n-1。所以，上述解决思路的时间复杂度就为 O(n)。
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSelect(nums []int, k int) (num int) {
	size := len(nums)
	pivot := rand.Intn(size)
	nums[size-1], nums[pivot] = nums[pivot], nums[size-1]

	pivot = size - 1
	left := 0
	right := size - 2
	for left <= right {
		if nums[left] <= nums[pivot] {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	nums[left], nums[pivot] = nums[pivot], nums[left]
	pivot = left

	//正好选择k
	if pivot == k-1 {
		num = nums[pivot]
	}
	// k 在右边
	if pivot < k-1 {
		num = quickSelect(nums[pivot+1:], k-pivot-1)
	}
	// k 在左边
	if pivot > k-1 {
		num = quickSelect(nums[:pivot], k)
	}
	return
}

func main() {
	rand.Seed(time.Now().Unix())
	nums := []int{3, 2, 4, 9, 7, 6, 5, 1, 8}
	k := 5

	num := quickSelect(nums, k)
	fmt.Printf("从小到大排序，第%d个元素为%d\n", k, num)
}
