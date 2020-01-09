/*
有 n 个城市通过 m 个航班连接。每个航班都从城市 u 开始，以价格 w 抵达 v。
现在给定所有的城市和航班，以及出发城市 src 和目的地 dst,你的任务是找到从 src 到 dst最多经过 k 站中转的最便宜价格线路。
如果没有这样的路线，则输出-1.
城市数 n 范围是[0-99]
航班数量范围是[0,n*(n-1)/2]
每个航班格式的格式是：（src,dst,price）
每个航班的价格范围是；[1,10000]
中转站数 k 范围是[0,n-1]
航班没有重复，且不存在环路

思路：这是一道典型的Dijkstra算法求最短便宜价格的问题，注意好约束条件 k 即可
*/

package main

import (
	"fmt"
)

const (
	maxCost         = 100000
	allNeedNodeNums = 1
)

var (
	graph     map[int]map[int]int
	costs     map[int]int
	parents   map[int]int
	processed []int

	dst int //目的地
	k   int //最多中转站数量
)

/* 自己写之前，没有抽象建模，没有建立数据结构，存储数据，
func findCheapestPrice(n, src, dst, k int, edges [][]int) int {

	var price []int = make([]int, n)
	//初始化从src到各个节点的价格
	for i := range price {
		if i == 0 {
			price[0] = 0
		}
		price[i] = 100000
	}
	fmt.Println(price)
	for _, v := range edges {
		if v[0] == 0 {
			price[v[1]] = v[2]
		}
	}

	//找个最便宜的节点，从它出发更新所连接的所有节点价格，循环这个过程，知道遍历了所有节点
	var cheaper []int
	p := 1

	for i := 1; i < n-1; i++ {
		if price[i] < price[p] {
			p = i
		}
	}
	cheaper = append(cheaper, p)
	for _, v := range edges {
		if v[0] == p {
			if price[p]+v[2] < price[v[1]] {
				price[v[1]] = price[p] + v[2]
			}
		}
	}
	if price[dst] != 100000 {
		return price[dst]
	}
	return -1
} */

func init() {
	//根据edges,先存储全图信息到map,第一张表
	graph = make(map[int]map[int]int, 3)
	g1 := make(map[int]int)
	g1[1] = 100
	graph[0] = g1

	g2 := make(map[int]int)
	g2[2] = 500
	graph[0] = g2

	g3 := make(map[int]int)
	g3[2] = 100
	graph[1] = g3

	//costs表，记录从出发点到所有点的距离，不能直接到达记录为无穷大，目前设为10000
	costs = make(map[int]int, 2)
	costs[1] = 100
	costs[2] = 500

	//parents表，记录到此节点最短路径的父节点,不存在则记录为-1
	parents = make(map[int]int, 2)
	parents[1] = 0
	parents[2] = 0

	dst = 2 //目的地初始化
	k = 0   //中转站 k 初始化
}

func findLowestCostNode() int {
	lowestCost := maxCost
	lowestCostNode := -1
	for k, v := range costs {
		if v < lowestCost && isNotProcessed(k) {
			lowestCost = v
			lowestCostNode = k
		}
	}
	return lowestCostNode
}

func isNotProcessed(node int) bool {
	for _, v := range processed {
		if v == node {
			return false
		}
	}
	return true
}

func findLowestCost() int {
	node := findLowestCostNode()
	for len(processed) != allNeedNodeNums && len(processed) < k {
		cost := costs[node]
		neighbors := graph[node]
		for k, c := range neighbors {
			newCost := cost + c
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode()
	}

	return costs[dst]
}

func main() {
	/* n := 3
	edges := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 1 */
	c := findLowestCost()
	if c == maxCost {
		fmt.Println("没有符合要求的飞行线路")
	} else {
		fmt.Println("存在符合条件的飞行线路，最低总价为：", c)
	}
}
