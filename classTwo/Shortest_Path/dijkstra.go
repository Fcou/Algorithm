package dijkstra

import (
	"container/heap"
	"container/list"
	"fmt"
)

const (
	INT_MAX = int(^uint(0) >> 1) //根据补码，有符号整数int最大值二进制表示，首位0，其余1
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
func (g *Graph) AddEdge(s, t, w int) { // s起点，t终点，w边的权重
	g.adj[s].PushBack(NewEdge(s, t, w)) // 存储顶点s指向的全部边，interface{} 各种类型都可以放入，取出时要断言
}

// 边的定义
type Edge struct {
	sid int // 边的起始顶点编号
	tid int // 边的终止顶点编号
	w   int // 权重
}

func NewEdge(sid, tid, w int) *Edge {
	return &Edge{sid, tid, w}
}

// 下面这个结构体是为了dijkstra实现用的
type Vertex struct {
	id   int // 顶点编号ID
	dist int // 从起始顶点到这个顶点的距离
}

// 从顶点s到顶点t的最短路径
func Dijkstra(s, t int, g *Graph) {
	predecessor := make([]int, g.v) // 用来还原最短路径,记录下标顶点的前驱顶点
	vertexes := make([]Vertex, g.v)
	for i := 0; i < g.v; i++ {
		vertexes[i] = Vertex{id: i, dist: INT_MAX} // 初始化全部顶点
	}
	// 优先队列初始化
	pq := make(PriorityQueue, g.v)
	for i := 0; i < len(vertexes); i++ {
		pq[i] = &Item{
			id:    vertexes[i].id,
			dist:  vertexes[i].dist,
			index: i,
		}
	}
	heap.Init(&pq)               // 初始化小顶堆
	inqueue := make([]bool, g.v) // 标记是否进入过队列,更新堆时有用

	vertexes[s].dist = 0 // 更新起点s的dist为0
	item := &Item{
		id:   vertexes[s].id,
		dist: vertexes[s].dist,
	}
	heap.Push(&pq, item) // 更新起点s的dist为0，不直接利用修改，没办法直接修改
	inqueue[s] = true
	//pq.update(item, item.id, 0)

	for pq.Len() > 0 {
		minVertex := heap.Pop(&pq).(*Item) // 取堆顶元素并删除
		if minVertex.id == t {             // 最短路径产生了
			break
		}
		for e := g.adj[minVertex.id].Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)          // 取出一条minVetex顶点相连的边
			nextVertex := vertexes[edge.tid] // minVertex-->nextVertex，获取min顶点指向的下一个next顶点的Vertex
			if minVertex.dist+edge.w < nextVertex.dist {
				nextVertex.dist = minVertex.dist + edge.w // 更新next的dist
				vertexes[nextVertex.id].dist = nextVertex.dist
				predecessor[nextVertex.id] = minVertex.id // 记录前驱顶点
				item := &Item{
					id:   nextVertex.id,
					dist: nextVertex.dist,
				}
				heap.Push(&pq, item) // 更新nextVertex.id的dist为 minVertex.dist + edge.w
				inqueue[nextVertex.id] = true
			}
		}
	}
	// 输出最短路径
	fmt.Printf("%d ", s)
	print(s, t, predecessor)
	for i := 0; i < len(vertexes); i++ {
		fmt.Printf("%d:%d\n", i, vertexes[i])
	}

}

func print(s, t int, predecessor []int) {
	if s == t {
		return
	}
	print(s, predecessor[t], predecessor)
	fmt.Printf("->%d ", t)
}
