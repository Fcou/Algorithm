/* 给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
示例 2:

输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
示例 3:

输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/same-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
package main

import "fmt"

// TreeNode is tree's node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

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

func InitTree(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Value: values[0]}
	if len(values) == 1 {
		return root
	}

	for _, v := range values[1:] {
		fmt.Println(v)
		node := &TreeNode{Value: v}
		InsertNodeToTree(root, node)
	}
	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Value == q.Value && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	treeOne := InitTree(1, 2, 2)
	treeTwo := InitTree(1, 2, 2)
	fmt.Println("Is this two trees same: ", isSameTree(treeOne, treeTwo))
}
