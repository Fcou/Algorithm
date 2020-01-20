/*
用广度优先算法，来解决走迷宫问题
*/

package main

import "fmt"

const (
	M      = 5 //M代表迷宫长
	N      = 5 //N代表迷宫宽
	Visted = 'v'
)

var (
	maze       [M][N]byte //'*'代表砖块，不可走；'o'代表空地，可以走
	source     point
	destinaton point
	P          myQueue
	parent     [M][N]point
)

//迷宫中的每个方格抽象成点，整个迷宫逻辑上就是点的集合,从左上角向下为x坐标轴，从左上角向右为y坐标轴
type point struct {
	x int
	y int
}

//记录每一步，当前点加方向，即可决定下个点的位置
type status struct {
	pt point
}

func findNextPoint(currrentPoint point, diretion int) (nextPoint point) {
	delta := [4]point{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
	}
	nextPoint.x = currrentPoint.x + delta[diretion].x
	nextPoint.y = currrentPoint.y + delta[diretion].y
	return
}

//根据status设计队列,链表结构,前出后进
type myQueue struct {
	first *myQueueNode
	end   *myQueueNode
}

type myQueueNode struct {
	val  status
	next *myQueueNode
}

func (queue *myQueue) init() {
	queue.first = nil
	queue.end = nil
}

func (queue *myQueue) Push(n status) {
	newNode := &myQueueNode{
		val:  n,
		next: nil,
	}
	if queue.first == nil && queue.end == nil {
		queue.first = newNode
		queue.end = newNode
		return
	}
	queue.end.next = newNode
	queue.end = newNode
}

func (queue *myQueue) Pop() (q status) {
	if queue.first == queue.end {
		q = queue.first.val
		queue.first = nil
		queue.end = nil
		return
	}
	q = queue.first.val
	queue.first = queue.first.next
	return
}

func (queue *myQueue) empty() bool {
	if queue.end == nil {
		return true
	}
	return false
}

func init() {
	//初始化迷宫，一旦其中的点被访问，修改对应的值为'v'，代表已访问
	maze = [M][N]byte{
		{'*', '*', '*', '*', '*'},
		{'*', 'o', '*', 'o', '*'},
		{'*', 'o', 'o', 'o', '*'},
		{'*', 'o', '*', 'o', '*'},
		{'*', '*', '*', '*', '*'},
	}
	source = point{x: 1, y: 1}
	destinaton = point{x: 3, y: 3}
	sourceStatus := status{pt: source} //从起点开始
	P.Push(sourceStatus)
}

//如果当前队列首不是出口，则弹出，同时将其领近可走的点都加入队列
func findDestinaton() bool {
	delta := [4]point{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
	}
	for !P.empty() {
		cur := P.Pop()
		if cur.pt.x == destinaton.x && cur.pt.y == destinaton.y {
			return true
		} else {
			for direction := 0; direction < 4; direction++ {
				neighbor := status{
					pt: point{
						x: (cur.pt.x + delta[direction].x),
						y: (cur.pt.y + delta[direction].y),
					},
				}
				if maze[neighbor.pt.x][neighbor.pt.y] != Visted && maze[neighbor.pt.x][neighbor.pt.y] != '*' {
					parent[neighbor.pt.x][neighbor.pt.y] = cur.pt
					maze[neighbor.pt.x][neighbor.pt.y] = Visted
					P.Push(neighbor)
				}
			}
		}
	}
	return false
}

func findParent() {
	fmt.Println("point:{", destinaton.x, destinaton.y, "}")
	tem := parent[destinaton.x][destinaton.y]
	for {
		fmt.Println("point:{", tem.x, tem.y, "}")
		tem = parent[tem.x][tem.y]
		if tem.x == source.x && tem.y == source.y {
			break
		}
	}
	fmt.Println("point:{", source.x, source.y, "}")
}
func main() {
	if findDestinaton() {
		findParent()
	}
}
