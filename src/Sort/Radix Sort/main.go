/*
基数排序，运用队列，按照从低位到高位依次执行一遍排序，最后得到正确顺序
*/
package main

import (
	"fmt"
	"math"
)

type myQueue struct {
	first *myQueueNode
	end   *myQueueNode
}

type myQueueNode struct {
	val  int
	next *myQueueNode
}

func (queue *myQueue) init() {
	queue.first = nil
	queue.end = nil
}

func (queue *myQueue) Push(n int) {
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

//刚开始这里设计错误，队列是先进先出，应该Pop队首元素
func (queue *myQueue) Pop() (q int, ok bool) {

	if queue.first == nil {
		ok = false
		return
	}
	q = queue.first.val
	ok = true
	queue.first = queue.first.next
	return
}

func main() {
	nums := []int{31, 99, 7, 56, 24, 38, 91, 17, 76, 39}
	radix := [10]myQueue{}

	for k := 0; k < 2; k++ {
		kk := int(math.Pow10(k))

		//第一次先按照个位,第二次按照十位，依次存入对应的队列中
		for _, v := range nums {
			radix[v/kk%10].Push(v)
		}
		//按照顺序从每个队列中依次取出
		nums = nil
		for i := 0; i < 10; i++ {
			for {
				q, ok := radix[i].Pop()
				if ok == false {
					break
				}
				nums = append(nums, q)
			}
		}
		radix = [10]myQueue{}
	}

	fmt.Println(nums)
}

//刚开始队列判断是否为空没有实现，代码冗余；个位和十位的处理应该进一步抽象统一，做成循环
//可以扩展，增加位数，也可以把其他排序问题，转换为数字，再运用基数排序思想排序
