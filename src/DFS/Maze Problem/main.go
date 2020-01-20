/*
用深度优先算法，来解决走迷宫问题
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
	P          myStack
)

//迷宫中的每个方格抽象成点，整个迷宫逻辑上就是点的集合,从左上角向下为x坐标轴，从左上角向右为y坐标轴
type point struct {
	x int
	y int
}

//记录每一步，当前点加方向，即可决定下个点的位置
type status struct {
	pt        point
	direction int //当前走的方向，初值为0，1，2，3代表可以走的四个方向，分别是3，6，9，12点钟方向，4代表无路可走
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

//根据status设计栈,链表结构
type myStack struct {
	first *myStackNode
	end   *myStackNode
}

type myStackNode struct {
	val  status
	next *myStackNode
}

func (stack *myStack) Push(n status) {
	newNode := &myStackNode{
		val:  n,
		next: nil,
	}
	if stack.first == nil && stack.end == nil {
		stack.first = newNode
		stack.end = newNode
		return
	}
	stack.end.next = newNode
	stack.end = newNode
}

func (stack *myStack) Pop() (q status) {

	if stack.empty() { //空栈返回false
		return
	}
	q = stack.end.val

	tem := stack.first
	if stack.first == stack.end {
		stack.first = nil
		stack.end = nil
	}
	for {
		if tem.next == stack.end {
			stack.end = tem
			break
		}
		tem = tem.next
	}
	return
}

func (stack *myStack) empty() bool {
	if stack.first == nil {
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
	sourceStatus := status{pt: source, direction: 0}
	P.Push(sourceStatus)
}

func findDestinaton() {
	for !P.empty() {
		cur := P.end //取出栈顶指针，并不弹出
		if cur.val.direction < 4 {
			nextPoint := findNextPoint(cur.val.pt, cur.val.direction)
			cur.val.direction++
			if maze[nextPoint.x][nextPoint.y] != Visted && maze[nextPoint.x][nextPoint.y] != '*' {
				maze[nextPoint.x][nextPoint.y] = Visted
				P.Push(status{pt: nextPoint, direction: 0})
				if nextPoint.x == destinaton.x && nextPoint.y == destinaton.y {
					break //找到终点，则跳出循环
				}
			}
		} else {
			P.Pop() //4个方向都走完了，则出栈
		}
	}
	s := P.first //依次从底向上，输出占栈中元素，就是迷宫路径，忽略最后出口的方向
	for s != nil {
		fmt.Println("point:", s.val.pt, "direction:", s.val.direction-1)
		s = s.next
	}
}

func main() {
	findDestinaton()
}
