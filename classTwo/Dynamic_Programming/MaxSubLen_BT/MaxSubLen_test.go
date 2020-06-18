package maxsublen

import (
	"fmt"
	"testing"
)

var MaxSubLen int = 0
var nums []int = []int{2, 9, 3, 6, 5, 1, 7}
var SubSequence []int = make([]int, 0)

//回溯算法
func findMaxSubSequence(SubSequence, nums []int, i int) {

	if i > len(nums)-1 { //遍历了原序列全部数字
		if MaxSubLen < len(SubSequence) {
			MaxSubLen = len(SubSequence)
		}
		return
	}

	findMaxSubSequence(SubSequence, nums, i+1) //不添加原序列第i位到子序列，不添加就不用考虑递增条件

	if len(SubSequence) == 0 || SubSequence[len(SubSequence)-1] < nums[i] { //剪枝，符合递增条件再考虑下一步
		SubSequence = append(SubSequence, nums[i]) //添加原序列第i位到子序列
		findMaxSubSequence(SubSequence, nums, i+1)
	}
}

func TestMSL(t *testing.T) {
	findMaxSubSequence(SubSequence, nums, 0)
	fmt.Println("MaxSubLen: ", MaxSubLen)
}
