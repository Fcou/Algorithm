/*
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同的值。这条路径可以经过也可以不经过根节点
两个节点之间的路径长度由他们之间的边数表示

分析：总路径=左边最长路径+右边最长路径 --递归条件  假设已经求出左右子树分别对应的最长路径，根据根结点值来判断计算总长
			顶点和左右节点都不一样  --基线条件

	从顶点开始遍历计算，比较，保留最长总路径
*/
package main

import "fmt"

//TreeNode 树节点定义
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

//logestUnivaluePath：求最长相同数值路径的主函数
func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root, &maxLen)
	return maxLen - 1
}

// 返回从 root 出发拥有相同 Value 值的线路上的 edge 数量
func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 0
	}

	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)
	res := 0

	// 左侧单边的最长距离
	if n.Left != nil && n.Value == n.Left.Value {
		*maxLen = max(*maxLen, l+1)
		res = max(res, l+1)
	}
	// 右侧单边的最长距离
	if n.Right != nil && n.Value == n.Right.Value {
		*maxLen = max(*maxLen, r+1)
		res = max(res, r+1)
	}
	// 通过根节点的最长边
	if n.Left != nil && n.Value == n.Left.Value &&
		n.Right != nil && n.Value == n.Right.Value {
		*maxLen = max(*maxLen, l+r+2)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//InitTree 将int数组初始化为树结构，返回根节点
func InitTree(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Value: values[0]}
	if len(values) == 1 {
		return root
	}

	for _, v := range values[1:] {
		node := &TreeNode{Value: v}
		InsertNodeToTree(root, node)
	}
	return root
}

//InsertNodeToTree 构建树的主要过程
func InsertNodeToTree(tree *TreeNode, node *TreeNode) {

	for tree != nil {

		if tree.Value > node.Value {
			if tree.Left == nil {
				tree.Left = node
				return
			} else {
				tree = tree.Left
			}

		} else {
			if tree.Right == nil {
				tree.Right = node
				return
			} else {
				tree = tree.Right
			}
		}

	}

}

func main() {
	tree := InitTree(1, 2, 2)
	longpath := longestUnivaluePath(tree)

	fmt.Println("The tree's the longest same path number is : ", longpath)
}
