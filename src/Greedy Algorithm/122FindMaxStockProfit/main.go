/*
给定一个数组，它的第i个元素是给定股票第i天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能多地交易。不能同时参与多笔交易，一次一笔。
例如：输入【7，1，5，3，6，4】输出 7

思路：首先判断是否是NP完全问题，因为设计到计算组合，时间复杂度为O(2^n),贪婪算法解决
碰到一个元素，就找后面比它大的最大元素，找不到就舍弃这个元素，找到就是一次交易，计算利润，
然后在此元素后重复这个过程，直到最后一个元素。
Ps.不应该是（后面比它大的最大元素），而是（后面比它大的第一个元素）
*/

package main

import "fmt"

var maxProfit int

func findMaxStockProfit(stockprice []int) {
	lens := len(stockprice)
	for i := 0; i < lens; i++ {
		buy := stockprice[i]
		key, sell := findLastMax(buy, stockprice[i:])
		if key == -1 {
			continue
		} else {
			maxProfit += sell - buy
			i = key + i
		}
	}
}

func findLastMax(b int, s []int) (key, value int) {
	max := b
	key = -1
	for k, v := range s {
		if v > max {
			max = v
			key = k
			value = v
			return
		}
	}
	return
}

//我写的太复杂，这个答案写的简单
func maxStockProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

func main() {
	stockprice := []int{7, 1, 5, 3, 6, 4}

	//findMaxStockProfit(stockprice)
	maxProfit = maxStockProfit(stockprice)
	fmt.Println("The max profit is :", maxProfit)
}
