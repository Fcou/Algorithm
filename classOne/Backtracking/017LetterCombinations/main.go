/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合

可以看成一种树，用广度、深度优先搜索也可以处理
*/

package main

import "fmt"

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

//letterCombinations 核心就是ret[j]+m[digits[i]][k]，之前字符串集合的每个组合+每一位数字代表的全部字母
func letterCombinations(digits string) []string {
	lens := len(digits)
	if lens == 0 {
		return nil
	}

	ret := []string{""}

	for i := 0; i < lens; i++ {
		temp := []string{}

		for j := 0; j < len(ret); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}
		ret = temp
	}
	return ret
}

func main() {
	digits := "234"

	combinations := letterCombinations(digits)

	for _, v := range combinations {
		fmt.Println(v)
	}

}
