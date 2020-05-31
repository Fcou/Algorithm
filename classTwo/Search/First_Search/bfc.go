// 之前写过基于迷宫问题的广度优先搜索和深度优先搜索，数据结构都采用的是，邻接矩阵
package First_Search

import (
	"container/list" //双向链表
	"fmt"
)

var graph *Graph

type Graph struct { // 无向图
	v   int          // 顶点的个数
	adj []*list.List // 邻接表,切片+链表，利用现有包
}

func Constructor(k int) *Graph { //初始化图
	g := &Graph{
		v:   k,
		adj: make([]*list.List, k),
	}
	for i := 0; i < k; i++ {
		g.adj[i] = list.New().Init()
	}
	return g
}

func (g *Graph) addEdge(s, t int) { // 无向图一条边存两次
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// 广度优先搜索（BFS），s 表示起点，t 表示终点
func bfs(s, t int) {
	if s == t {
		return
	}
	visited := make([]bool, graph.v) //记录已访问的节点，节点编号是从0到graph.v，没有重复
	visited[s] = true
	queue := make([]int, 0)      // 队列,注意初始化大小要为0
	queue = append(queue, s)     // s 入队
	prev := make([]int, graph.v) // prev 用来记录搜索路径
	for i := 0; i < graph.v; i++ {
		prev[i] = -1
	}
	for len(queue) != 0 { //队列不为空，则一直遍历
		w := queue[0]                                           // 取出首元素
		queue = queue[1:]                                       // 首元素出队
		for e := graph.adj[w].Front(); e != nil; e = e.Next() { // 开始遍历当前节点的相邻节点，注意遍历写法
			q := e.Value.(int) // 取出 w 相连接的节点
			if !visited[q] {   // 没被访问
				prev[q] = w // prev[q]存储的是，顶点 q 是从前驱顶点 w 遍历过来的。
				if q == t { // 如果找到目标，则打印路径返回
					print(prev, s, t)
					return
				}
				visited[q] = true        // 记录已被访问
				queue = append(queue, q) // 新节点入队
			}
		}
	}
	fmt.Println("No exists")
}

func print(prev []int, s, t int) { // 递归打印s->t的路径
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t]) //prev[t] 表示 t 的前驱顶点
	}
	fmt.Printf("-%d-", t)
}
