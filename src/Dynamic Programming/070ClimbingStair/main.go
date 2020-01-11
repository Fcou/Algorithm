/*
假设你正在爬楼梯，每次你可以爬1或2个台阶，总台阶为n,问你有多少种不同的方法爬到楼顶
*/

package main

import "fmt"

func climbingStair(n int) (m int) {
	if n <= 1 {
		m = 1
		return
	}

	m = climbingStair(n-1) + climbingStair(n-2)

	return
}

func main() {
	n := 20

	m := climbingStair(n)

	fmt.Println(n, "个台阶爬楼方法共有：", m)
}
