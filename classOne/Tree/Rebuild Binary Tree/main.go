/*
在已知前序遍历、中序遍历二叉树的结果后（数组存储），重新建立二叉树结构
*/

package main

import "fmt"

var (
	preNums = []int{2, 0, 1, 3, 4, 5, 6, 7}
	inNums  = []int{0, 1, 2, 3, 4, 5, 6, 7}
)

type tnode struct {
	data  int
	left  *tnode
	right *tnode
	//parent *tnode
}

//n为要处理数组的长度范围，确保n>0
func rebulidByInAndPre(pre, in []int, root []tnode, n int) (ok bool) {
	pivot := findPivot(pre[0], in)
	if pivot == -1 {
		return false
	}

	leftSize := pivot - 0
	rightSize := 0 + n - pivot - 1

	if leftSize == 0 {
		root[0].left = nil
	} else if !rebulidByInAndPre(pre[1:], in[0:], root[1:], leftSize) {
		return false
	} else {
		root[0].left = &root[1]
	}

	if rightSize == 0 {
		root[0].right = nil
	} else if !rebulidByInAndPre(pre[1+leftSize:], in[1+leftSize:], root[1+leftSize:], rightSize) {
		return false
	} else {
		root[0].right = &root[1+leftSize]
	}
	root[0].data = pre[0]

	return true
}

func findPivot(root int, inNums []int) int {
	for k, v := range inNums {
		if root == v {
			return k
		}
	}
	return -1
}

func main() {
	tree := make([]tnode, 8) //存放树，布局为：根，左子树，右子树，方便找到根的位置,前序形式
	ok := rebulidByInAndPre(preNums, inNums, tree, 8)
	if ok {
		fmt.Println("重新构建二叉树成功")
		for _, v := range tree {
			fmt.Println(v)
		}
	} else {
		fmt.Println("重新构建二叉树失败")
	}
}
