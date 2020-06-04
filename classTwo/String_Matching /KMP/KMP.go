package main

import (
	"fmt"
)

func getNexts(pattern string) []int {
	m := len(pattern)       // 模式串长度
	nexts := make([]int, m) //如果 next[i-1]=k-1，也就是说，子串 b[0, k-1]是 b[0, i-1]的最长可匹配前缀子串。
	for index := range nexts {
		nexts[index] = -1 // 初始化为-1
	}
	//假设已知nexts[i-1],利用nexts[i-1]求nexts[i],类似动态规划思想
	for i := 1; i < m-1; i++ {
		j := nexts[i-1] // pattern[0,j] 是pattern[0, i-1]的最长可匹配前缀子串

		for pattern[j+1] != pattern[i] && j >= 0 { // 下一位不相等,这个时候就不能简单地通过 next[i-1]得到 next[i]了
			//以pattern[i]为最后一个相同字符，找前面最长前缀子串，并且pattern[j+1] == pattern[i]，否则一直循环查找
			//也就是在pattern[0, i-1]的最长可匹配前缀子串，找到次长最长可匹配前缀子串
			j = nexts[j] // 求出pattern[0, j]的最长可匹配前缀子串，
		}

		if pattern[j+1] == pattern[i] { // 下一位相等，否则跳不出上面循环，或者j==-1
			j++ // pattern[0,j+1] 是pattern[0, i]的最长可匹配前缀子串
		}

		nexts[i] = j
	}

	return nexts
}

func findByKMP(s string, pattern string) int {
	n := len(s)
	m := len(pattern)
	if n < m {
		return -1
	}

	nexts := getNexts(pattern) // nexts 数组，确定出现坏字符后，模式串第几位和主串现有位比较

	j := 0                   //记录模式串中当前匹配位置
	for i := 0; i < n; i++ { //记录主串中当前匹配位置
		for j > 0 && s[i] != pattern[j] { // 模式串中第j位与主串当前第i位不匹配，一直不匹配则一直循环
			j = nexts[j-1] + 1 // 根据nexts数组计算出，模式串[0,j-1]的最长可匹配前缀子串+1位开始和主串i比较
		}

		if s[i] == pattern[j] { //一定相等，不相等则跳不出上面的循环,或在j==0
			if j == m-1 { // 已完全匹配
				return i - m + 1
			}
			j++ //没完全匹配但该位相同，j后移一位继续判断,i也会后移一位
		}
	}

	return -1
}

func main() {
	s := "abc abcdab abcdabcdabde"
	pattern := "bcdabd"
	fmt.Println(findByKMP(s, pattern)) //16

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "abab"
	fmt.Println(findByKMP(s, pattern)) //11

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "ababacd"
	fmt.Println(findByKMP(s, pattern)) //-1

	s = "hello"
	pattern = "ll"
	fmt.Println(findByKMP(s, pattern)) //2
}
