package main

import "fmt"

//SelectionSort 选择排序，每次选出剩下元素中最小的，排在最前面
func SelectionSort(nums []int) {
	lens := len(nums)
	for i := 0; i < lens; i++ {
		temp := i
		for j := i + 1; j < lens; j++ {
			if nums[j] < nums[temp] {
				temp = j
			}
		}
		nums[i], nums[temp] = nums[temp], nums[i]

	}

}

func main() {
	nums := []int{3, 6, 1, 9, 2, 4, 8, 7, 5}

	SelectionSort(nums)

	fmt.Println(nums)
}
