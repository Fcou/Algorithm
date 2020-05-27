package SkipList

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

//跳表节点结构体
type skipListNode struct {
	//跳表保存的值，实际存储数值
	v interface{}
	//用于排序的分值，有序链表排序
	score int
	//层高，随机得来
	level int
	//每层前进指针，类似多层Next指针
	forwards []*skipListNode
}

//新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
}

//跳表结构体
type SkipList struct {
	//跳表头节点，头节点是层次最高的
	head *skipListNode
	//跳表当前最大层数
	level int
	//跳表长度，总节点个数
	length int
}

//实例化跳表对象
func NewSkipList() *SkipList {
	//头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL) // MinInt32  = -1 << 31
	return &SkipList{head, 1, 0}
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

//插入节点到跳表中,结果有三种情况: 空节点则不插入返回1，节点已存在不插入则返回2，成功插入则返回0
func (sl *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找插入位置
	cur := sl.head
	//记录每层的路径
	update := [MAX_LEVEL]*skipListNode{}
	//从最高层开始遍历
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		//每层遍历到最后为nil
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			//score用来标记链表从小到大顺序？
			if cur.forwards[i].score > score {
				update[i] = cur //存储该层要向下的节点指针，即路径节点
				break
			}
			//前进一步，指向下一个链表节点
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	//原有节点连接，根据level跟其之前的每个节点都要建立连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

//查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if nil == v || sl.length == 0 {
		return nil
	}

	cur := sl.head
	//从最高层开始遍历
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

//删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找前驱节点
	cur := sl.head
	//记录前驱路径
	update := [MAX_LEVEL]*skipListNode{}
	//要从最高层每层遍历，找到每层的前驱节点
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	//每层的前驱节点都要指向要删除节点的下一个节点
	for i := cur.level - 1; i >= 0; i-- {
		//更新跳表当前最大层数
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}
