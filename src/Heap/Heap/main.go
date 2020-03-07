//看看container/heap的源码，写法值得我学习
//堆接口组合了 sort.Interface, 而sort.Interface，需要实现三个方法：
//Len() int /   Less(i, j int) bool  /  Swap(i, j int)
//再加上堆接口定义的两个方法：Push(x interface{})   /  Pop() interface{}。
//只要实现了上面这五个方法，定义了一个最小堆。

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap provides heap operations for any type that implements
// heap.Interface. A heap is a tree with the property that each node is the
// minimum-valued node in its subtree.
//
// The minimum element in the tree is the root, at index 0.
//
// A heap is a common way to implement a priority queue. To build a priority
// queue, implement the Heap interface with the (negative) priority as the
// ordering for the Less method, so Push adds items while Pop removes the
// highest-priority item from the queue. The Examples include such an
// implementation; the file example_pq_test.go has the complete source.
//
package heap

import "sort"

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// Init has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface interface { //需要自己实现接口，Push()和Pop()
	sort.Interface      //匿名嵌套
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n) //效率比 依次加入建堆算法 高
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	h.Push(x)        //添加到最后一位
	up(h, h.Len()-1) //最后一位上浮
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)   //最小元素换到最后一位,最后位置元素换到第0位
	down(h, 0, n)  //第0位下沉，注意它不会换到n,最多第n-1位
	return h.Pop() //弹出最后一位，即最小元
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)        //把i换到最后一位
		if !down(h, i, n) { //把最后一位执行下沉操作，执行过下沉操作则说明最后一位大于目前孩子，下沉后堆恢复正常。
			up(h, i) //否则，说明最后一位不大于目前孩子，需要执行上浮操作。
		}
	}
	return h.Pop() //把i换到最后一位，并更新堆，最后弹出最后一位
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) { //h.Len()，这样下沉范围包括最后一位;不下沉，就上浮。即不小于孩子，就大于孩子。
		up(h, i)
	}
}

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j) //如果父节点下标为j，比i小，则交换对应元素
		j = i        //下标i赋值为j,继续判断是否还要上浮
	}
}

func down(h Interface, i0, n int) bool { //不包括n
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j) //先找到i的左右孩子二者中最小的，下标为j，如果比i小，则交换对应元素
		i = j        //下标i赋值为j,继续判断是否还要下沉
	}
	return i > i0 //如果执行过下沉，则下标i一定会大于i0，返回true
}
