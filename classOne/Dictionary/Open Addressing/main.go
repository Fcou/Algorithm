/* 在链接法中，如果不同键值却将有相同的映射值，即有不同键值的元素却映射到散列表中的同一位置，
那么就采用链表的方法，将映射到同一位置的元素插入到同一个链表之中，当需要删除， 查询元素时，只需要遍历该链表即可，
链接法在最坏情况下删除和查询元素的时间代价为O(n)O(n)
今天我们来讲散列表中另外一种解决冲突的方法，那就是开放定址法(open addressing)。

假如你在外面旅游时，吃坏东西，急需上厕所，当你好不容易找到一件洗手间的时候，发现排了好多人，这时你会怎么做？

如果是链接法:排队不就行了，我就在外面等，迟早会排到我的
如果是开放定址法:直接放弃现有厕所，去寻找新的厕所
没错，放弃已被占用的位置，寻找新的插入位置就是开放定址法的思想，开放定址法中的开放二字指的是没有被占用的位置，
定址指的是确定位置。开放定址法中，所有的元素都放在散列表中(链接法放在链表中)。也就是说散列表中的每一个位置，
要么有元素，要么没有元素。当需要删除，查询元素时，我们从某一个位置开始，按照某种特定的确定下一个位置的方法来
检查所有表项，直到找到目标元素，或者没有找到。
————————————————
版权声明：本文为CSDN博主「qeesung」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/ii1245712564/article/details/46846197 */

package main

import (
	"fmt"
)

//pnode 节点
type pnode struct {
	key   int  //存储信息
	empty bool //标记是否为空,true为空
	gone  bool //标记是否被删除，true为已被删
}

//linearProbing 线性探查函数,返回查询到的数组位置。k 为关键字，offset 为第几次探查，arraySize 为数组长度
func linearProbing(k int, offset int, arraySize int) int {
	return (hashKey(k, arraySize) + offset) % arraySize
}

func hashKey(k int, arraySize int) int {
	return k % arraySize
}

//init  数组初始化
func initArray(arr []pnode) {
	for k := range arr {
		arr[k].empty = true
		arr[k].gone = false
	}
}

//insert 插入操作，先查询，以防存入重复元素。每一位都要查询，全部都失败则说明没位置插入。已满要重散列，再插入。
func insert(arr []pnode, k int) []pnode {
	size := len(arr)

	if find(arr, k) == size { //不重复

		for i := 0; i < size; i++ { //开始插入
			pos := linearProbing(k, i, size)

			if arr[pos].empty {
				arr[pos].key = k
				arr[pos].empty = false
				return arr
			} else if arr[pos].gone {
				arr[pos].key = k
				arr[pos].gone = false
				return arr
			}
		}

		newArr := doubleHashSize(arr) //已满，则重散列
		insert(newArr, k)             //重新插入k
		return newArr
	}
	return arr
}

//find 查询操作,查询到则返回位置，如果没查到则返回数组长度
func find(arr []pnode, k int) int {
	size := len(arr)

	for i := 0; i < size; i++ { //每一位都要查询
		pos := linearProbing(k, i, size)
		if arr[pos].empty {
			return size //后面也不会有
		} else if !arr[pos].gone && arr[pos].key == k {
			return pos
		}
	}
	return size
}

//findKey 查询Key操作，对find()封装一下，输出些消息
func findKey(arr []pnode, k int) {
	size := len(arr)
	if find(arr, k) == size {
		fmt.Println("目前散列表中不存在该元素:", k)
	} else {
		fmt.Println("目前散列表中存在该元素:", k)
	}
}

//delete 删除操作,查询到则gone变成true,代表被删除，保留k值，返回true；否则，返回false
func delete(arr []pnode, k int) bool {
	pos := find(arr, k)
	if pos == len(arr) {
		fmt.Println("目前散列表中不存在该元素:", k)
		return false
	}
	arr[pos].gone = true
	fmt.Println("目前散列表中已删除该元素:", k)
	return true
}

//doubleHashSize 重散列，当哈希表都被填满or大于装填因子，翻倍容量，重新散列
func doubleHashSize(arr []pnode) []pnode {
	W := make([]pnode, len(arr)*2)

	initArray(W)

	for k := range arr {
		if !arr[k].empty && !arr[k].gone {
			insert(W, arr[k].key)
		}
	}
	return W
}

func main() {
	//H := new([5]pnode) //散列数组

	H := make([]pnode, 5)

	initArray(H)

	H = insert(H, 8) //切片重散列后，注意 H 要重新赋值
	H = insert(H, 9)
	H = insert(H, 111)
	H = insert(H, 1)
	H = insert(H, 8)
	H = insert(H, 88)
	H = insert(H, 100)
	H = insert(H, 200)
	H = insert(H, 33)
	H = insert(H, 9898)
	H = insert(H, 33335)
	H = insert(H, 2)

	delete(H, 44)
	delete(H, 100)

	findKey(H, 100)
	findKey(H, 2)
	for k := range H {
		fmt.Println(H[k])
	}
}
