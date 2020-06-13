/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组，求出这个数组中的逆序对的总数。

我的思考:
0.描述了一种情况，求出现这种情况的次数
1.遍历一遍每个数字，后面有比自己小的，逆序对的总数+1
2.递归,还是要遍历，没有提高效率
3.贪心，动态，类似都不适用
4.先排序，排序完成，对比每个元素的位置变化，直接计算逆序对总数

递归都是分治思想的运用，但一般情况，递归每次缩小范围有限，例如n运用n-1的结果，每次只缩小1
高效的分治，每次问题都能缩小一半，例如归并排序，二叉搜索
在排序分治的过程中，处理一些问题，也是高效的
*/
package main

import "fmt"

func reversePairs(nums []int) int {
	var sum int
	if len(nums) == 0 {
		return sum
	}
	//合并两个数组，两路归并，从小到大排序
	var merge func(left []int, right []int) []int
	merge = func(left []int, right []int) []int {
		var result []int
		for len(left) != 0 && len(right) != 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:] //改变切片起始位置，从而不用改变归并的下标，妙啊
			} else {
				result = append(result, right[0])
				//其他代码就是正常的归并排序，这行记录每次合并时，left中的每个元素对right中每个元素产生的逆序对
				//如果，left[0]>right[0],则left[0]及其之后的所有元素都比right[0]大，因为left已从小到大排好序
				//一次性计算出，现有left数组全部元素相对于right[0]产生的逆序对总数
				//每次归并，累加逆序对总数，最终为全部结果
				//前面的元素，会跟所有的后面元素进行比较、归并，所以会计算出全部逆序对
				sum += len(left)
				//只增加了以上一行代码
				right = right[1:]
			}
		}

		for len(left) != 0 {
			result = append(result, left[0])
			//sum += len(left)   right中已没有比left小的元素，不会产生逆序对
			left = left[1:]
		}

		for len(right) != 0 {
			result = append(result, right[0])
			right = right[1:]
		}

		return result
	}
	var mergeSort func(nums []int) []int
	mergeSort = func(nums []int) []int {
		lens := len(nums)
		if lens < 2 {
			return nums
		}
		//划分成左右两个数组
		middle := lens / 2
		left := nums[0:middle]
		right := nums[middle:]
		return merge(mergeSort(left), mergeSort(right)) //假设已对左右子树归并排序完，再对两个数组进行合并，即可完成
	}

	mergeSort(nums)
	return sum
}

func main() {
	nums := []int{8, 4, 5, 7, 10}
	fmt.Println(reversePairs(nums))
}
