package Binary_Tree

var (
	proportion = 0.5 //根据这个比例计算数组中根节点的位置，根在数组总长度0.5的位置
	nums       = []int{0, 1, 2, 3, 4, 5, 6, 7}
)

type BinarySearchTree struct {
	head *Node
}

type Node struct {
	data   int
	left   *Node
	right  *Node
	parent *Node //如果没有父节点，查找前后节点比较麻烦
}

// 根据数组和比例，来建立二叉搜索树
func BSTGeneration(start int, size int) (rootTree *Node) {
	if size < 1 {
		rootTree = nil
		return
	}
	rootTree = new(Node) //特别注意引用类型的变量要初始化，分配内存空间，否则无法存储变量；值类型不用，系统已默认分配内存
	pivot := int(float64(size-1) * proportion)
	LSize := pivot
	RSize := size - 1 - pivot
	root := start + pivot

	leftSubtree := BSTGeneration(start, LSize)
	rightSubtree := BSTGeneration(start+pivot+1, RSize)

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

// 二叉搜索树查找操作
func (bst *BinarySearchTree) Find(x int) *Node {
	p := bst.head
	for p != nil {
		if p.data == x {
			return p
		} else if p.data < x {
			p = p.left
		} else {
			p = p.right
		}
	}
	return nil
}

// 二叉搜索树插入操作
func (bst *BinarySearchTree) Insert(x int) {
	p := bst.head
	if p == nil {
		p = &Node{data: x}
		bst.head = p
		return
	}
	for p != nil {
		if p.data == x { //如果已有相同元素，则不插入
			return
		} else if p.data < x {
			if p.left == nil {
				p.left = &Node{data: x，parent: p,}
				return
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &Node{data: x, parent: p,}
				return
			}
			p = p.right
		}
	}
	return
}

// 二叉搜索树删除操作
func (bst *BinarySearchTree) Delete(x int) {
	p := bst.head // p指向要删除的节点，初始化指向根节点
	for p != nil {
		if p.data == x {
			break
		} else if p.data < x {
			p = p.left
		} else {
			p = p.right
		}
	}
	if p == nil {
		return // 没有找到要删除的节点
	}

	// 要删除的节点有两个子节点
	if p.left != nil && p.right != nil { // 查找右子树中最小节点，来替代要删除的节点
		minP = p.right
		for minP != nil {
			minP = minP.left
		}
		// 找到后，交换data,下面就变成了删除之前的minP节点了
		p.data = minP.data
		p = minP
	}

	// 删除节点是叶子节点或者仅有一个子节点,找到其子节点
	var child *Node // p的子节点
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}
	// 开始正式删除操作
	if p.parent == nil {
		bst.head = child // 删除的是根节点
	} else if p.parent.left == p {
		p.parent.left = child
	} else {
		p.parent.right = child
	}
}

// 二叉搜索树查找最大值操作
func (bst *BinarySearchTree) FindMax() *Node {
	p := bst.head
	if p == nil {
		return nil
	}
	return rightMost(p)
}

//以p为根，找到它的最大元
func rightMost(p *Node) *Node {
	for p.right != nil {
		p = p.right
	}
	return p
}

// 二叉搜索树查找最小值操作
func (bst *BinarySearchTree) FindMin() *Node {
	p := bst.head
	if p == nil {
		return nil
	}
	return leftMost(p)
}

//以p为根，找到它的最小元
func leftMost(p *Node) *Node {
	for p.left != nil {
		p = p.left
	}
	return p
}

// 二叉搜索树查找前驱节点操作
func (bst *BinarySearchTree) PreviousNode() *Node {
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

// 二叉搜索树查找后继节点操作
func (bst *BinarySearchTree) NextNode(p *Node) *Node {
	if p.right != nil { //如果p的右子树不为空，则下一元素就是右子树的最小元
		return leftMost(p.right)
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
