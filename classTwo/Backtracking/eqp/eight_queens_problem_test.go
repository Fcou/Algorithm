package eqp

import (
	"fmt"
	"testing"
)

var result []int = make([]int, 8) //全局或成员变量,下标表示行,值表示queen存储在哪一列
var count int = 0                 //记录全部摆放次数

func cal8queens(row int) {
	if row == 8 { // 8个棋子都放置好了，打印结果
		printQueens(result)
		return // 8行棋子都放好了，已经没法再往下递归了，所以就return
	}
	for column := 0; column < 8; column++ { // 每一行都有8种放法
		if isOk(row, column) { // 如果放法满足要求
			result[row] = column // 第row行的棋子放到了column列
			cal8queens(row + 1)  // 可以放，***则考察下一行***
		}
	}
	// 如果每行的8种方法都不行，则会返回上一级，上一行重新考虑下一列的摆放
}

//判断row行column列放置是否合适
func isOk(row, column int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- { // 逐行往上考察每一行
		if result[i] == column { // 第i行的column列有棋子吗？
			return false
		}
		if leftup >= 0 { //考察左上对角线：第i行leftup列有棋子吗？
			if result[i] == leftup {
				return false
			}
		}
		if rightup < 8 { // 考察右上对角线：第i行rightup列有棋子吗？
			if result[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true //由于每行只放一个，所以不必检查行，列和两个对角线检查了即可
}

// 打印出一个二维矩阵
func printQueens(result []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	count++
}

func TestEQP(t *testing.T) {

	cal8queens(0) // 这将打印出所有符合条件的摆放

	fmt.Println("all count:", count)
}
