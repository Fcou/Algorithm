/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

*/
package main

import "fmt"

// 二分查找法求平方根
func mySqrt(x float64) float64 {
	var low, high, ans float64
	low, high, ans = 0, x, -1
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid <= x {
			ans = mid
			low = mid + 0.000001
		} else {
			high = mid - 0.000001
		}
	}
	return ans
}

func mySqrtTwo(x int) int {
	low, high := 0, x
	ans := -1
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid <= x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(mySqrt(99))
}
