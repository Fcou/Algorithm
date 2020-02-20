/*
桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
为了使桶排序更加高效，我们需要做到这两点：
1.在额外空间充足的情况下，尽量增大桶的数量
2.使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中
同时，对于桶中元素的排序，选择何种比较排序算法对于性能的影响至关重要。
运行时间O(n),不一定比直接sort快，内存碎片多
*/
package main

import "fmt"

//桶内排序
func sortInBucket(bucket []int) { //元素个数少，此处用插入排序方式，可以用任意其他排序方式
	length := len(bucket)

	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if bucket[i] < bucket[j] {
				bucket[i], bucket[j] = bucket[j], bucket[i]
				i = j
			} else {
				break
			}
		}

	}

}

//获取数组最大值
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// bucketSort 桶排序,arr为要排序的数组，e为每个桶中元素期望个数
func bucketSort(arr []int, e int) {
	n := len(arr)                 //数组元素个数
	num := n / e                  //桶数
	max := getMaxInArr(arr) + 1   //max（数组最大值+1）
	buckets := make([][]int, num) //创建桶空间

	//分配入桶，将数组元素变为【0，1）,乘以桶数就是要放入的桶编号
	for i := 0; i < len(arr); i++ {
		index := arr[i] * num / max //分配桶编号
		buckets[index] = append(buckets[index], arr[i])
	}

	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(arr[tmpPos:], buckets[i])

			tmpPos += bucketLen
		}
	}

}

func main() {
	arr := []int{10, 3, 8, 20, 23, 12, 7, 5, 19, 17}
	bucketSort(arr, 2)
	fmt.Println(arr)
}
