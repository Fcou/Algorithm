/*
给定一个字符串S和一个字符C。返回一个代表字符串S中每个字符到字符串S中的字符C的最短距离的数组。
第一种想法：从左到右依次求出到第一个C的最短距离，再从右到左再求依次，记录两次每个字符的最短距离。
第二种想法：把字符串转变为根节点为C的数组，求每个叶子节点到根的高度
这两种想法好像跟 并查集 ，都没关系
*/
package main

import "fmt"

//使用双下标，一次循环，同时比较从左和从右计算的距离，取二者最小值。
//left,只影响第一个C之后的元素；right，只影响最后一个C之前的元素；
//第一个C和最后一个C之间的距离计算两次比较，两头只用计算一次。这个算法设计的巧妙！
func shortestToChar(S string, C byte) []int {
	n := len(S)
	res := make([]int, n)
	for i := range res {
		res[i] = n
	}

	left, right := -n, 2*n //这样取值，保证程序开始计算最左边元素到left的位置为n，最右边距离right标记的位置为n

	for i := 0; i < n; i++ {
		j := n - i - 1
		if S[i] == C {
			left = i
		}
		if S[j] == C {
			right = j
		}
		res[i] = min(res[i], dist(i, left))
		res[j] = min(res[j], dist(j, right))
	}

	return res
}

//比较从左或从右距离标记字符C的最小距离
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//计算距离标记字符C的距离
func dist(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func main() {
	S := "aeel"
	var C byte = 'e'

	lon := shortestToChar(S, C)
	fmt.Println(lon)
}
