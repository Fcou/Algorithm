/*
基数排序，运用队列，按照从低位到高位依次执行一遍排序，最后得到正确顺序
这里按照每位来排序的排序算法要是稳定的，否则这个实现思路就是不正确的。
因为如果是非稳定排序算法，那最后一次排序只会考虑最高位的大小顺序，完全不管其他位的大小关系，那么低位的排序就完全没有意义了。
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

		//第一次先按照个位,第二次按照十位，依次存入对应的队列中,这里保证稳定性
		for _, v := range nums {
			radix[v/kk%10].Push(v)
		}
		//按照顺序从每个队列中依次取出，重新存入nums切片中
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
		radix = [10]myQueue{} //队列清空，下次再利用
	}

	fmt.Println(nums)
}

//可以扩展，增加位数，也可以把其他排序问题，转换为数字，再运用基数排序思想排序
//实际上，有时候要排序的数据并不都是等长的，比如我们排序牛津字典中的 20 万个英文单词，最短的只有 1 个字母，最长的我特意去查了下，
//有 45 个字母，中文翻译是尘肺病。对于这种不等长的数据，基数排序还适用吗？
//实际上，我们可以把所有的单词补齐到相同长度，位数不够的可以在后面补“0”，因为根据ASCII 值，所有字母都大于“0”，
//所以补“0”不会影响到原有的大小顺序。这样就可以继续用基数排序了。
