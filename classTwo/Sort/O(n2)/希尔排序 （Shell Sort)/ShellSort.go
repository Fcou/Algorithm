/*
希尔排序(Shell's Sort)是插入排序的一种又称“缩小增量排序”（Diminishing Increment Sort），是直接插入排序算法的一种更高效的改进版本。
希尔排序是非稳定排序算法。
希尔排序是把记录按下标的一定增量分组，对每组使用直接插入排序算法排序；
随着增量逐渐减少，每组包含的关键词越来越多，当增量减至1时，整个文件恰被分成一组，算法便终止。
*/

package main

import "fmt"

// 插入排序，a表示数组，n表示下标间隔数
func insertionSort(a []int, n int) {
	// 分成n组，分别执行插入排序
	for m := 0; m < n+1; m++ {
		// 间隔为n,执行插入排序
		for i := m + n; i < len(a); i = i + n {
			value := a[i]
			j := i - n
			// 查找插入的位置
			for ; j >= m; j = j - n {
				if a[j] > value {
					a[j+n] = a[j]
				} else {
					break // a[j] <= value  之前的数都小
				}
			}
			a[j+n] = value // 插入
		}
	}
}

func shellSort(a []int, n int) {
	size := n / 2
	for size != 0 {
		insertionSort(a, size)
		size = size / 2
	}

}

func main() {
	nums := []int{2, 4, 77, 3, 99, 23, 4, 18, 1, 8, 42, 14, 56}
	shellSort(nums, len(nums))
	fmt.Println(nums)
}
