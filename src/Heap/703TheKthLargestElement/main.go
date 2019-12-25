/* 设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。

你的 KthLargest 类需要一个同时接收整数 K 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素。

示例:

int k = 3;
int[] arr = [4,5,8,2,3,5,10,9];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
说明:
你可以假设 nums 的长度≥ k-1 且k ≥ 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */


package main


import (
	"container/heap"
	"fmt"
   )


   type intHeap []int

   //KthLargest object will be instantiated and called as such: obj := Constructor(k, nums);
   // param_1 := obj.Add(val);
   type KthLargest struct {
	k    int
	ih intHeap
   }

   
   
   // Constructor 创建 KthLargest
   func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)
   
	for len(h) > k {
	 heap.Pop(&h)
	}
   
	return KthLargest{
	 k:  k,
	 ih: h,
	}
   }
   
   // Add 负责添加元素
   func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.ih, val)
   
	if len(kl.ih) > kl.k {
	 heap.Pop(&kl.ih)
	}
   
	return kl.ih[0]
   }
   
   
   
   func (h intHeap) Len() int {
	return len(h)
   }
   
   func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
   }
   
   func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
   }
   func (h *intHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
   }
   
   func (h *intHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
   }



   func main(){
	
	arr := [] int{4,5,8,2,3,5,10,9}
	kl := Constructor(4,arr)
	fmt.Println(kl)

	fmt.Println( kl.Add(3),kl )
	fmt.Println( kl.Add(5) ,kl)
	fmt.Println( kl.Add(111),kl )
	fmt.Println( kl.Add(9),kl )
	fmt.Println( kl.Add(4) ,kl)

   }
   
   //题目的意思就是要求数组从大到小第K大的元素，K初始化定义好，之后固定
   //Add()方法是插入新元素，通过计算继续返回第K大元素
   //维护一个前K大元素的数据结构就行，这里选用最小堆，用数组存储
   //如果比最小堆堆顶小，则舍去新元素，返回堆顶。如果比堆顶大，则删除最小堆，插入新元素，返回堆顶