//来源：https://zhuanlan.zhihu.com/p/34989978
// 散列 + 双向链表，利用golang现有包
// https://time.geekbang.org/column/article/64858 文章里设计更加底层、基础，实现都要自己写 散列表、链表、各种操作代码
package lru

import "container/list" //双向链表

// type LRUCache struct {
// 	capacity int
// 	cache    map[int]*Pair //之前觉得这样更简洁，其实缺少功能，不可取
// }

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element //存储指向链表元素的指针，查询方便
	list     *list.List            //链表用来确定先后顺序，确定该淘汰删除哪个
}

//其中链表的元素中存储的Value 就是Pair结构体，接口太方便了
type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok { //map结构搜索是否存有key
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if this.list.Len() >= this.capacity {
			delete(this.cache, this.list.Back().Value.(Pair).key) //map删除操作
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	}
}
