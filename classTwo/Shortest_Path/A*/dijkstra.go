package dijkstra

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
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
	f    int // 新增：f(i)=g(i)+h(i)
	x    int // 新增：顶点在地图中的坐标（x, y）
	y    int // 新增：顶点在地图中的坐标（x, y）
}

func (v *Vertex) setVertex(id, x, y int) {
	v.id = id
	v.dist = INT_MAX
	v.f = INT_MAX
	v.x = x
	v.y = y
}

// // 新增一个方法，添加顶点的坐标
// func (v *Vertex) addVertex(id, x, y int) {
// 	vertexes[id] = new Vertex(id, x, y)
// }

// 从顶点s到顶点t的最短路径
func Dijkstra(s, t int, g *Graph) {
	predecessor := make([]int, g.v) // 用来还原最短路径,记录下标顶点的前驱顶点
	vertexes := make([]Vertex, g.v)
	vertexes[0].setVertex(0, 0, 1) // 初始化全部顶点
	vertexes[1].setVertex(1, 3, 2)
	vertexes[2].setVertex(2, 1, 4)
	vertexes[3].setVertex(3, 2, 5)

	// 优先队列初始化
	pq := make(PriorityQueue, g.v)
	for i := 0; i < len(vertexes); i++ {
		pq[i] = &Item{
			id:    vertexes[i].id,
			f:     vertexes[i].f, // dist替换为f，f(i)=dist+hManhattan(i,t),按照f构建小顶堆
			index: i,
		}
	}
	heap.Init(&pq) // 初始化小顶堆

	vertexes[s].dist = 0 // 更新起点s的dist为0
	vertexes[s].f = 0
	item := &Item{
		id:   vertexes[s].id,
		dist: vertexes[s].f,
	}
	heap.Push(&pq, item) // 更新起点s的f为0，不直接利用修改，没办法直接修改

	for {
		minVertex := heap.Pop(&pq).(*Item) // 取堆顶元素并删除
		if minVertex.id == t {             // 最短路径产生了
			break
		}
		for e := g.adj[minVertex.id].Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)          // 取出一条minVetex顶点相连的边
			nextVertex := vertexes[edge.tid] // minVertex-->nextVertex，获取min顶点指向的下一个next顶点的Vertex
			if minVertex.dist+edge.w < nextVertex.dist {
				nextVertex.dist = minVertex.dist + edge.w
				nextVertex.f = nextVertex.dist + hManhattan(nextVertex, vertexes[t])
				vertexes[nextVertex.id].dist = nextVertex.dist // 更新next的dist
				vertexes[nextVertex.id].f = nextVertex.f       // 更新next的f
				predecessor[nextVertex.id] = minVertex.id      // 记录前驱顶点
				item := &Item{
					id: nextVertex.id,
					f:  nextVertex.f,
				}
				heap.Push(&pq, item) // 更新堆，nextVertex.id的dist为 minVertex.dist + edge.w
			}
		}
	}
	// 输出最短路径
	fmt.Printf("%d ", s)
	print(s, t, predecessor)

	// 输出每个顶点到s顶点的最短距离
	fmt.Println()
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

func hManhattan(v1, v2 Vertex) int { // Vertex表示顶点
	return int(math.Abs(float64(v1.x-v2.x)) + math.Abs(float64(v1.y-v2.y)))
}
