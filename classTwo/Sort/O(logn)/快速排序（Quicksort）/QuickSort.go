package main

import (
	"fmt"
	"math/rand"
	"time"
)

//快速排序算法，原地置换版本，不需要额外的内存空间
func quickSort(nums []int) {
	size := len(nums)
	if size < 2 {
		return
	}
	pivot := rand.Intn(size)
	nums[size-1], nums[pivot] = nums[pivot], nums[size-1] //pivot先存在最后一位

	pivot = size - 1
	left := 0
	right := size - 2
	// left 都小于等于 pivot，右边都大于 pivot
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

	quickSort(nums[:pivot])
	quickSort(nums[pivot+1:])
}

func main() {
	rand.Seed(time.Now().Unix()) //随机初始化
	nums := []int{3, 2, 4, 9, 7, 6, 5, 1, 8, 99, 0}

	quickSort(nums)
	fmt.Println(nums)
}
