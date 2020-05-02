/*
合并查找：
1.之前写过Disjoint set算法，利用数组下标指代父亲节点，两者原理一致
2.复杂数据都是基础数据类型组成，类似 图 由 点 组成，这个是基础
3.顶层设计，过程拆分清楚，一个Find，一个Union
4.利用struct组合封装 多个有关系元素
*/
package main

import "fmt"

//PR 每个节点的基础类型，由他们的数组表示一堆点的集合
type PR struct{     
	parent int      //表示该节点的父亲节点位置，由于是静态的，用数组下标表示即可，存储成int类型即可
	rank   int   //表示秩，节点合并的图形可以做成树形，秩代表：树的高度和树中节点数（单调递增），用于平衡
}

//UnionDS 合并x和y两个根节点，先利用findDS找到对应的两个根
func UnionDS(DS []PR, x int, y int)  {
	if x == y {
		return
	}

	if DS[x].rank < DS[y].rank {
		DS[x].parent = y
	} else if DS[y].rank < DS[x].rank{
		DS[y].parent = x
	} else {
		DS[y].parent = x
		DS[x].rank++
	}
}

//findDS 查找x元素根节点
func findDS(DS []PR, x int) (root int){
	root = x
	for DS[root].parent != root{
		root = DS[root].parent    //root最终停在x的根节点
	}
	for x != root {
		par := DS[x].parent
		DS[x].parent = root       //路径压缩，路径上经过的节点全指向根
		x = par              
	}
	return
}

func main() {
	var DS []PR = make([]PR, 7)

	for i := 0; i < 7; i++ {
		DS[i].parent = i
		DS[i].rank = 0
	}

	UnionDS(DS,findDS(DS, 1),findDS(DS, 5))
	UnionDS(DS,findDS(DS, 1),findDS(DS, 6))
	UnionDS(DS,findDS(DS, 2),findDS(DS, 4))
	UnionDS(DS,findDS(DS, 5),findDS(DS, 4))
	
	for i := 0; i < 7; i++ {
		fmt.Printf(" %d",i)
	}
	fmt.Printf("\n")
	for _, v := range DS {
		fmt.Printf(" %d",v.parent)
	}
	fmt.Printf("\n")
	for _, v := range DS {
		fmt.Printf(" %d",v.rank)
	}
}