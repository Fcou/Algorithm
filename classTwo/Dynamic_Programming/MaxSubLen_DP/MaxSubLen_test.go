package maxstatelen

import (
	"fmt"
	"testing"
)

//动态规划算法
func findMaxstateSequence(nums []int, n int) int {

	state := make([]int, n)
	state[0] = 1 // 代表遍历到nums第0位，最长递增子串长度为1

	var maxLen int
	for i := 1; i < n; i++ {
		maxLen = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				maxLen = maxTwo(maxLen, state[j]+1) //a[0...i] 的最长子序列为: a[i] 之前所有比它小的元素中子序列长度最大的 + 1
			}
		}
		state[i] = maxLen
	}

	maxLen = -1
	// 筛选以一个统一标准，测算出的长度集合，求出最大值
	for i := 0; i < n; i++ {
		fmt.Println(state[i])
		if maxLen < state[i] {
			maxLen = state[i]
		}
	}

	return maxLen
}

func maxTwo(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestMSL(t *testing.T) {
	var nums []int = []int{2, 9, 3, 6, 5, 1, 7}

	fmt.Println("MaxstateLen: ", findMaxstateSequence(nums, len(nums)))
}
