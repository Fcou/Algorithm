/*
直方图排序，解决桶排序内存碎片多而产生的改进算法
*/

package main

import "fmt"

func histogramSort(arr []int, maxValue int, e int) []int {
	arrSize := len(arr)
	num := arrSize / e //桶个数
	max := getMaxInArr(arr) + 1

	C := make([]int, num+1) // 计数数组
	B := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		C[arr[i]*num/max+1]++
	}
	fmt.Println("C:", C)
	for i := 1; i < num+1; i++ {
		C[i] += C[i-1]
	}
	fmt.Println("C:", C)

	for _, v := range arr {
		j := C[v*num/max]
		B[j] = v
		C[v*num/max] = C[v*num/max] + 1
	}
	fmt.Println("C:", C)
	fmt.Println("B:", B)
	left := 0
	for i := 0; i < num-1; i++ {
		right := left + C[i]
		sortInBucket(B[left:right])
		left = right
	}
	return B
}

//getMaxInArr 获取数组最大值
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//sortInBucket 桶内排序
func sortInBucket(bucket []int) { //元素个数少，此处用插入排序方式，可以用任意其他排序方式
	length := len(bucket)

	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if bucket[i] < bucket[j] {
				bucket[i], bucket[j] = bucket[j], bucket[i]
				i = j
			} else {
				break
			}
		}

	}

}
func main() {
	nums := []int{5, 6, 9, 3, 4, 1, 2, 7, 8, 10}
	nums = histogramSort(nums, 9, 5)
	fmt.Println("nums:", nums)
}
