//给定一个数组，给定比例，来构造一个二叉树

package main

import "fmt"

var (
	proportion = 0.3 //根据这个比例计算数组中根节点的位置，根在数组总长度0.3的位置
	nums       = []int{0, 1, 2, 3, 4, 5, 6, 7}
)

type tnode struct {
	data   int
	left   *tnode
	right  *tnode
	parent *tnode
}

func treeGeneration(start int, size int) (rootTree *tnode) {
	if size < 1 {
		rootTree = nil
		return
	}
	rootTree = new(tnode) //特别注意引用类型的变量要初始化，分配内存空间，否则无法存储变量；值类型不用，系统已默认分配内存
	pivot := int(float64(size-1) * proportion)
	LSize := pivot
	RSize := size - 1 - pivot
	root := start + pivot

	leftSubtree := treeGeneration(start, LSize)
	rightSubtree := treeGeneration(start+pivot+1, RSize)

	rootTree.data = nums[root]
	rootTree.left = leftSubtree
	rootTree.right = rightSubtree
	if leftSubtree != nil {
		leftSubtree.parent = rootTree
	}
	if rightSubtree != nil {
		rightSubtree.parent = rootTree
	}

	return
}

func main() {
	root := treeGeneration(0, len(nums))
	root.parent = nil
	fmt.Println(root)
}
