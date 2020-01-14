/* 设计你的循环队列实现。 循环队列是一种线性数据结构，
其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

循环队列的一个好处是我们可以利用这个队列之前用过的空间。
在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。
但是使用循环队列，我们能使用这些空间去存储新的值。

你的实现应该支持如下操作：

MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。


示例：

MyCircularQueue circularQueue = new MycircularQueue(3); // 设置长度为 3

circularQueue.enQueue(1);  // 返回 true

circularQueue.enQueue(2);  // 返回 true

circularQueue.enQueue(3);  // 返回 true

circularQueue.enQueue(4);  // 返回 false，队列已满

circularQueue.Rear();  // 返回 3

circularQueue.isFull();  // 返回 true

circularQueue.deQueue();  // 返回 true

circularQueue.enQueue(4);  // 返回 true

circularQueue.Rear();  // 返回 4



提示：

所有的值都在 0 至 1000 的范围内；
操作数将在 1 至 1000 的范围内；
请不要使用内置的队列库。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-circular-queue
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

/* package main

// MyCircularQueue 结构体
type MyCircularQueue struct {
	queue []int
	k     int
}

// Constructor initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, 0, k*3),
		k:     k,
	}
}

// EnQueue insert an element into the circular queue. Return true if the operation is successful.
func (m *MyCircularQueue) EnQueue(value int) bool {
	if len(m.queue) == m.k {
		return false
	}
	m.queue = append(m.queue, value)
	return true
}

// DeQueue delete an element from the circular queue. Return true if the operation is successful.
func (m *MyCircularQueue) DeQueue() bool {
	if len(m.queue) == 0 {
		return false
	}

	m.queue = m.queue[1:]
	return true
}

// Front get the front item from the queue.
func (m *MyCircularQueue) Front() int {
	if len(m.queue) == 0 {
		return -1
	}

	return m.queue[0]
}

// Rear get the last item from the queue. */
/* func (m *MyCircularQueue) Rear() int {
	if len(m.queue) == 0 {
		return -1
	}
	return m.queue[len(m.queue)-1]
}

// IsEmpty checks whether the circular queue is empty or not.
func (m *MyCircularQueue) IsEmpty() bool {
	return len(m.queue) == 0
}

// IsFull checks whether the circular queue is full or not.
func (m *MyCircularQueue) IsFull() bool {
	return len(m.queue) == m.k
} */

/**
* Your MyCircularQueue object will be instantiated and called as such:
* obj := Constructor(k);
* param_1 := obj.EnQueue(value);
* param_2 := obj.DeQueue();
* param_3 := obj.Front();
* param_4 := obj.Rear();
* param_5 := obj.IsEmpty();
* param_6 := obj.IsFull();
 */

//这只是一个普通队列，不是循环队列 */
package main

var (
	nums []int
	lens = 10
	f    = 0
	r    = 0
)

func init() {
	nums = make([]int, lens)
}

func full() bool {
	return f == (r+1)%lens
}

func empity() bool {
	return f == r
}
func pop() (x int, ok bool) {
	if empity() {
		ok = false
	}
	x = nums[f]
	f = (f + 1) % lens
	ok = true
	return
}
func push(x int, ok bool) {
	if full() {
		ok = false
		return
	}
	nums[r] = x
	r = (r + 1) % lens
	ok = true
}

//循环意义下的左闭右开区间，[f,r),少用一个空间，装不满
//代码优化，%取模慢，代替方法：如果 ++w， w == lens, w = 0;
//规避取模运算，判断条件 f== (r+1)%n 替换为  (r+1) < lens? f == (r+1); f==0
//本质：抽屉原理，固定r,f有lens种取值，表达lens种状态，但其实有lens+1种状态
//这样就不会一一对应，两种取值对应一个取值就会引起冲突
//以上方法，就是少装一个元素，来给判断空or满留足情况
//或者利用一个计数器来，利用额外信息存储工具 记录队列全部信息表达
