/*
把字符串数组G,存入M和IM中，形成双向映射
*/
package main

import "fmt"

func main() {
	var G []string = []string{"O(1)", "O(n)", "O(logn)", "O(n^2)", "O(n^n)", "O(n)"}

	var M []string //等价于var M map[int]string
	var IM map[string]int = make(map[string]int)

	for _, v := range G {
		_, ok := IM[v] //检查IM中已存在字符串v
		if ok {
			continue //已存在，则跳过
		} else {
			M = append(M, v)   //向M中添加新元素
			IM[v] = len(M) - 1 //向IM中添加新元素
		}
	}

	fmt.Println(M)
	fmt.Println(IM)
}
