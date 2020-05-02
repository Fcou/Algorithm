/*
给定一个二叉树，检查它是否是镜像对称
思路1：就是要检查节点顺序，可以先把树的节点顺序存下来，运用双端队列，一层一层依次加入，弹出，最后不留元素说明是镜像对称
需要用双端队列，同时从前和后，遍历对比每层节点
golang 没有内置一些基础数据结构，麻烦

思路2：递归，基准条件：单一节点对称，左右节点不同时是空节点不对称；递归条件：左子树根节点的值等于右子树根节点的值，
左子树左节点=右子树右节点，同时左子树右节点=右子树左节点
*/

package main

// TreeNode is tree's node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recur(root.Left, root.Right)
}

func recur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && recur(left.Left, right.Right) && recur(left.Right, right.Left)
}
