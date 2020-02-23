/* 背包问题 */

package main

import "fmt"

type commodity struct {
	name   string
	price  int
	weight int
}

// type knapsack struct {
// 	maxCapacity     int
// 	currentCapacity int
// 	currentPrice    int
// }

func maxTwoNum(a, b int) (x int, ok bool) {
	if a > b {
		x = a
		ok = true
		return
	} else {
		x = b
		ok = false
		return
	}
}

func main() {
	var c [3]commodity //商品列表，定义名称，价格，重量，下标代表依次计算判断的顺序
	c[0] = commodity{name: "guitar", price: 1500, weight: 1}
	c[1] = commodity{name: "sound", price: 3000, weight: 4}
	c[2] = commodity{name: "laptop", price: 2000, weight: 3}

	totalPrice := [3][4]int{}         //初始化总价格矩阵
	chooseCommodity := [3][4]string{} //初始化组成最高价的商品矩阵
	ok := true

	for i := 0; i < len(totalPrice); i++ {
		for j := 0; j < len(totalPrice[i]); j++ {
			if j+1 >= c[i].weight && i == 0 {
				totalPrice[i][j] = c[i].price
				chooseCommodity[i][j] = c[i].name
			} else if j+1 < c[i].weight {
				totalPrice[i][j] = totalPrice[i-1][j]
				chooseCommodity[i][j] = chooseCommodity[i-1][j]
			} else if j+1 == c[i].weight {
				totalPrice[i][j], ok = maxTwoNum(totalPrice[i-1][j], c[i].price)
				if ok {
					chooseCommodity[i][j] = chooseCommodity[i-1][j]
				} else {
					chooseCommodity[i][j] = c[i].name
				}

			} else if j+1 > c[i].weight {
				totalPrice[i][j], ok = maxTwoNum(totalPrice[i-1][j], c[i].price+totalPrice[i][j-c[i].weight])
				if ok {
					chooseCommodity[i][j] = chooseCommodity[i-1][j]
				} else {
					chooseCommodity[i][j] = c[i].name + " And " + chooseCommodity[i][j-c[i].weight]
				}
			}
		}
	}
	for i := 0; i < len(totalPrice); i++ {
		fmt.Println(totalPrice[i])
	}
	for i := 0; i < len(totalPrice); i++ {
		fmt.Println(chooseCommodity[i])
	}
	fmt.Println("背包里最多能装商品的总价为：", totalPrice[2][3])
	fmt.Println("总价最高商品组合为：", chooseCommodity[2][3])
}
