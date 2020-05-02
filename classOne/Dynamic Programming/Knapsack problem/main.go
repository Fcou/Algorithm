/*
背包问题
坑1：二维切片初始化
*/

package main

import "fmt"

type commodity struct {
	name   string
	price  int
	weight int
}

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

func findMaxPrice(c []commodity, totalPrice [][]int, chooseCommodity [][]string) {
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
				totalPrice[i][j], ok = maxTwoNum(totalPrice[i-1][j], c[i].price+totalPrice[i-1][j-c[i].weight])
				if ok {
					chooseCommodity[i][j] = chooseCommodity[i-1][j]
				} else {
					chooseCommodity[i][j] = c[i].name + " And " + chooseCommodity[i-1][j-c[i].weight]
				}
			}
		}
	}
}

func main() {
	var c []commodity //初始化商品列表，定义名称，价格，重量，下标代表依次计算判断的顺序
	c = append(c, commodity{name: "guitar", price: 1500, weight: 1})
	c = append(c, commodity{name: "sound", price: 3000, weight: 4})
	c = append(c, commodity{name: "laptop", price: 2000, weight: 3})

	totalPrice := make([][]int, 3) //初始化总价格3*4矩阵
	for i := 0; i < len(totalPrice); i++ {
		totalPrice[i] = make([]int, 4)
	}
	chooseCommodity := make([][]string, 3) //初始化组成最高价的商品3*4矩阵
	for i := 0; i < len(totalPrice); i++ {
		chooseCommodity[i] = make([]string, 4)
	}

	findMaxPrice(c, totalPrice, chooseCommodity)

	for i := 0; i < len(totalPrice); i++ {
		fmt.Println(totalPrice[i])
	}
	for i := 0; i < len(totalPrice); i++ {
		fmt.Println(chooseCommodity[i])
	}
	fmt.Println("背包里最多能装商品的总价为：", totalPrice[2][3])
	fmt.Println("总价最高商品组合为：", chooseCommodity[2][3])
}
