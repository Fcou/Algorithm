//插入排序
package main

import "fmt"

func insertSort(nums []int) { //元素个数少，可以用插入排序方式
	length := len(nums)

	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				i = j
			} else {
				break
			}
		}
	}
}

func main() {
	nums := []int{4, 7, 1, 5, 9, 2, 8}
	insertSort(nums)
	fmt.Println(nums)
}
