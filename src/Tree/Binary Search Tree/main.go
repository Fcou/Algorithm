//先构建一个二叉搜索树，对二叉搜索树求最大元，最小元，每个节点的上一个值，下一个值
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

//以下为二叉搜索树的相关操作
//以p为根，找到它的最小元
func leftMost(p *tnode) *tnode {
	for p.left != nil {
		p = p.left
	}
	return p
}

//以p为根，找到它的最大元
func rightMost(p *tnode) *tnode {
	for p.right != nil {
		p = p.right
	}
	return p
}

//找到p节点的，下一元素；如果没有则返回 nil
func nextNode(p *tnode) *tnode {
	if p.right != nil {
		return leftMost(p.right) //如果p的右子树不为空，则下一元素就是右子树的最小元
	}

	last := p //如果p的右子树为空,则p的下一元素就是以p为最大元树根的父节点
	for p != nil {
		if last == p.left {
			break
		}
		last = p
		p = p.parent
	}
	return p
}

//找到p节点的，上一元素
func previousNode(p *tnode) *tnode {
	if p.left != nil {
		return rightMost(p.left) //如果p的左子树不为空，则下一元素就是左子树的最大元
	}

	last := p //如果p的左子树为空,则p的上一元素就是以p为最小元树根的父节点
	for p != nil {
		if last == p.right {
			break
		}
		last = p
		p = p.parent
	}
	return p
}

//遍历操作,从小到大
func traversalTree(root *tnode) {
	//p初始化为最小元，利用nextNode函数，依次寻找下一元素
	for p := leftMost(root); p != nil; p = nextNode(p) {
		fmt.Println(p.data)
	}
}

//插入操作，在查找失败的位置插入
func insertNode(root, p *tnode) {
	for root != nil {
		if p.data > root.data {
			root = root.right
		} else {
			root = root.left
		}
	}
	root = p
}

//删除操作，用该元素的上一个/下一个元素替换它
func deleteNode(root, p *tnode) {
	for root != nil {
		if p.data > root.data {
			root = root.right
		} else if p.data < root.data {
			root = root.left
		} else {
			d := nextNode(root) //root目前就是要删除的节点，可以找到它的下一个节点，替换它
			if d != nil {
				d.left = root.left
				d.right = root.right
				d.parent = root.parent
				root = d
			} else {
				d := previousNode(root) //如果没有下一个节点，找前一个节点替换
				if d != nil {
					d.left = root.left
					d.right = root.right
					d.parent = root.parent
					root = d
				} else {
					root = nil //前后都没有节点，说明只有一个节点，直接设置为空即可
				}
			}

		}
	}

}

func main() {
	root := treeGeneration(0, len(nums))
	root.parent = nil
	fmt.Println(root)
}
