/* 假设我们有一个 n 乘以 n 的矩阵 w[n][n]。矩阵存储的都是正整数。棋子起始位置在左上角，终止位置在右下角。
我们将棋子从左上角移动到右下角。每次只能向右或者向下移动一位。从左上角到右下角，会有很多不同的路径可以走。
我们把每条路径经过的数字加起来看作路径的长度。那从左上角移动到右下角的最短路径长度是多少呢？ */

package minDist

func minDistDP(matrix [][]int, n int) int {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, n)
	}
	sum := 0
	for j := 0; j < n; j++ { // 初始化states的第一行数据
		sum += matrix[0][j]
		states[0][j] = sum
	}

	sum = 0
	for i := 0; i < n; i++ { // 初始化states的第一列数据
		sum += matrix[i][0]
		states[i][0] = sum
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			states[i][j] = matrix[i][j] + minTwo(states[i][j-1], states[i-1][j])
		}
	}

	return states[n-1][n-1]
}

func minTwo(x, y int) int {
	if x < y {
		return x
	}
	return y
}
