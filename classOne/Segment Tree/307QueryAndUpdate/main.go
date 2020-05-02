/*
	给定一个整数数组nums，求出数组从索引 i 到 j (i<=j)范围内元素的总和，包含i，j两点。
	update(i,val) 函数可以通过将下标为i的数值更新为val，从而对数列进行修改。
*/

package main

import "fmt"

var (
	MaxLen int = 1000
)

//tree为要建立的数组结构，node为tree当前要建立的数组元素下标，需要利用nums数组和 start到end 的元素
func buildTree(tree []int, node int, nums []int, start int, end int) {
	if start == end {
		tree[node] = nums[start]
	} else {
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		mid := (start + end) / 2

		buildTree(tree, leftNode, nums, start, mid)
		buildTree(tree, rightNode, nums, mid+1, end)
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}

//updateTree修改nums数组中下标为idx的数组元素的值为val,之后更新对应字段数tree
func updateTree(tree []int, node int, nums []int, start int, end int, idx int, val int) {
	if start == end {
		nums[idx] = val
		tree[node] = val
	} else {

		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		if idx >= start && idx <= mid {
			updateTree(tree, leftNode, nums, start, mid, idx, val)

		} else {
			updateTree(tree, rightNode, nums, mid+1, end, idx, val)
		}
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}

//queryTree给定数组nums的左右边界L，R，求边界内数组元素之和
func queryTree(tree []int, node int, nums []int, start int, end int, L int, R int) int {
	if end < L || start > R {
		return 0
	} else if start == end {
		return tree[node]
	} else if L <= start && R >= end {
		return tree[node]
	} else {

		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2

		sumLeft := queryTree(tree, leftNode, nums, start, mid, L, R)
		sumRight := queryTree(tree, rightNode, nums, mid+1, end, L, R)

		return sumLeft + sumRight
	}
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	lens := len(nums)
	var tree []int = make([]int, 20, MaxLen)

	buildTree(tree, 0, nums, 0, lens-1)
	fmt.Println(tree)

	updateTree(tree, 0, nums, 0, lens-1, 4, 6)
	fmt.Println(tree)

	sum := queryTree(tree, 0, nums, 0, lens-1, 2, 5)
	fmt.Println(sum)

}
