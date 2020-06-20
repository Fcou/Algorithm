package dfs

import (
	"container/list"
	"fmt"
)

// 图的定义
type Graph struct {
	v   int          // 顶点的个数
	adj []*list.List // 邻接表,注意从0开始
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

// 拓扑排序使用DFS算法
func TopoSortByDFS(g *Graph) {
	// 先构建逆邻接表，边s->t表示，s依赖于t，t先于s
	inverseAdj := make([]*list.List, g.v)
	for i := 0; i < g.v; i++ {
		inverseAdj[i] = list.New()
	}
	// 通过邻接表生成逆邻接表
	for i := 0; i < g.v; i++ {
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			w := e.Value.(int)        // i->w
			inverseAdj[w].PushBack(i) // w->i
		}
	}
	visited := make([]bool, g.v) //记录顶点是否被访问
	// 深度优先遍历图
	for i := 0; i < g.v; i++ {
		if visited[i] == false {
			visited[i] = true
			dfs(i, inverseAdj, visited)
		}
	}
}

// 用递归写的深度优先遍历
func dfs(vertex int, inverseAdj []*list.List, visited []bool) {
	for e := inverseAdj[vertex].Front(); e != nil; e = e.Next() {
		w := e.Value.(int)
		if visited[w] == true {
			continue
		}
		visited[w] = true
		dfs(w, inverseAdj, visited)
	}
	fmt.Printf("->%d", vertex) //  先把vertex这个顶点可达的所有顶点都打印出来之后，再打印它自己
}
