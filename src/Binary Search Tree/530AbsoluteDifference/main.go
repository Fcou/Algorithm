/*
给定一个所有节点为非负值的二叉搜索树，求树中任意两个节点的差的绝对值的最小值。（树中最少两个节点）
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type state struct {
	minDiff, previous int
}

func getMinimumDifference(root *TreeNode) int {
	st := state{1024, 1024}
	search(root, &st)
	return st.minDiff
}

// NOTICE: BST 的递归遍历方法
//我第一想法是:遍历找出最大值和最小值，然后计算
//这里算法的意思是，遍历每个节点算出差值，通过比较差值大小，留下最小的
func search(root *TreeNode, st *state) {
	if root == nil {
		return
	}

	search(root.Left, st)

	newDiff := diff(st.previous, root.Val)
	if st.minDiff > newDiff {
		st.minDiff = newDiff
	}

	st.previous = root.Val

	search(root.Right, st)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
