package main

import "fmt"

//Vertices是图中全部顶点数,Parent数组是图中顶点对应的父节点关系数组，Rank数组记录对应顶点在聚合过程中树的高度
var (
	Vertices int = 6
	Parent       = make([]int, Vertices)
	Rank         = make([]int, Vertices)
)

//初始化自动调用，父节点Parent都为-1，高度rank都为0
func init() {
	for i := 0; i < Vertices; i++ {
		Parent[i] = -1
		Rank[i] = 0
	}
}

//寻找下标x元素的根节点
func findRoot(x int, parent []int) int {
	xRoot := x
	for parent[xRoot] != -1 {
		xRoot = parent[xRoot]
	}
	return xRoot
}

//合并x,y两个点，组成集合（其实是树结构.两个点合并成功，返回true；两个点已经在一个集合则不用合并，返回false
func unionTwoVertices(x int, y int, parent []int, rank []int) bool {
	xRoot := findRoot(x, parent)
	yRoot := findRoot(y, parent)

	if xRoot == yRoot {
		return false
	} else {
		//通过比较两树的高度，来判断如何合并根节点，降低合并后树的高度
		if rank[xRoot] > rank[yRoot] {
			parent[yRoot] = xRoot
		} else if rank[yRoot] > rank[xRoot] {
			parent[xRoot] = yRoot
		} else {
			parent[xRoot] = yRoot
			rank[yRoot]++
		}
		return true
	}

}

func main() {

	edges := [6][2]int{
		{0, 1}, {1, 2}, {1, 3},
		{2, 4}, {3, 4}, {2, 5},
	}

	for i := 0; i < 6; i++ {
		x := edges[i][0]
		y := edges[i][1]
		if unionTwoVertices(x, y, Parent, Rank) == false {
			fmt.Println("该图存在闭环")
			return
		}
	}
	fmt.Println("该图不存在闭环")
}
