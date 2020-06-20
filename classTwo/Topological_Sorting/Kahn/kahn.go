package kahn

import (
	"container/list"
	"fmt"
)

// 图的定义
type Graph struct {
	v   int          // 顶点的个数
	adj []*list.List // 邻接表,用于存储下标数字节点指向的全部其他数字，由于数量不确定，用链表结构动态添加
}

// 初始化并生成v个顶点的一个图对象
func InitGraph(v int) *Graph {
	g := &Graph{
		v:   v,
		adj: make([]*list.List, v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = list.New()
	}
	return g
}

// 添加图的边，s、t 为对应的顶点数字
func (g *Graph) AddEdge(s, t int) { // s先于t，边s->t
	g.adj[s].PushBack(t) // 存储顶点s指向的全部边，interface{} 各种类型都可以放入，取出时要断言
}

// 拓扑排序使用Kahn算法
func TopoSortByKahn(g *Graph) {
	inDegree := make([]int, g.v) // 统计每个顶点的入度
	for i := 0; i < g.v; i++ {
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			inDegree[e.Value.(int)]++ // i->w
		}
	}

	queue := list.New() // 存放最终符合拓扑排序的节点序列
	for i := 0; i < g.v; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i) // 放入最开始入度为 0 的节点
		}
	}
	for queue.Len() != 0 {
		i := queue.Back().Value.(int)                       // 返回队尾元素
		queue.Remove(queue.Back())                          // 删除队尾元素
		fmt.Printf("->%d", i)                               // 打印
		for e := g.adj[i].Front(); e != nil; e = e.Next() { // 将i节点指向的全部节点，入度都减1
			k := e.Value.(int)
			inDegree[k]--
			if inDegree[k] == 0 {
				queue.PushBack(k) // 过程中检查入度为 0 的节点，放入队列中
			}
		}
	}
}
