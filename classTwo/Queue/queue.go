/*
队列最大的特点就是先进先出，主要的两个操作是入队和出队。
跟栈一样，它既可以用数组来实现，也可以用链表来实现。用数组实现的叫顺序队列，用链表实现的叫链式队列。
特别是长得像一个环的循环队列。
在数组实现队列的时候，会有数据搬移操作，要想解决数据搬移的问题，我们就需要像环一样的循环队列。

应用：
生产者-消费者模型 阻塞队列：
阻塞队列其实就是在队列基础上增加了阻塞操作。简单来说，就是在队列为空的时候，从队头取数据会被阻塞。
因为此时还没有数据可取，直到队列中有了数据才能返回；
如果队列已经满了，那么插入数据的操作就会被阻塞，直到队列中有空闲位置后再插入数据，然后再返回。

并发队列：线程安全的队列
最简单直接的实现方式是直接在 enqueue()、dequeue() 方法上加锁，
但是锁粒度大并发度会比较低，同一时刻仅允许一个存或者取操作。实际上，基于数组的循环队列，
利用 CAS 原子操作，可以实现非常高效的并发队列。这也是循环队列比链式队列应用更加广泛的原因。

有限资源排队
实际上，对于大部分资源有限的场景，当没有空闲资源时，基本上都可以通过“队列”这种数据结构来实现请求排队。
*/

//一个可扩展队列，基于切片（可自动扩展）
package queue

type myQueue struct {
	items []string
	count int //当前队列内元素个数
}

func (this *myQueue) newQueue(n int) *myQueue {
	this.items = make([]string, n)
	this.count = 0
	return this
}

func (this *myQueue) push(item string) {
	this.items = append(this.items, item)
	this.count++
	return
}

func (this *myQueue) pop() (result string, ok bool) {
	if this.count < 1 {
		result = ""
		ok = false
		return
	}
	result = this.items[0]
	ok = true
	this.count--
	this.items = this.items[1:]
	return
}

//循环队列
type CircularQueue struct {
	items []string
	n     int //队列容量
	head  int
	tail  int
}

func (this *CircularQueue) newCircularQueue(n int) *CircularQueue {
	this.items = make([]string, n)
	this.n = n
	this.head = 0
	this.tail = 0
	return this
}

func (this *CircularQueue) push(item string) bool {
	if (this.tail+1)%this.n == this.head {
		return false
	}
	this.items[this.tail] = item
	this.tail = (this.tail + 1) % this.n
	return true
}

func (this *CircularQueue) pop() (result string, ok bool) {
	if this.head == this.tail {
		result = ""
		ok = false
		return
	}
	result = this.items[this.head]
	ok = true
	this.head = (this.head + 1) % this.n
	return
}
