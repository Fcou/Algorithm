package main

import (
	"fmt"
	"math"
)

//bc: pattern char index hash mapping  坏字符规则，快速定位在模式串中最后出现位置
func generateBC(pattern string) []int {

	bc := make([]int, 256) // 数组的下标对应字符的 ASCII 码值

	for index := range bc {
		bc[index] = -1
	}

	for index, char := range pattern {
		bc[int(char)] = index // 将模式串中的每个字符及其下标都存到散列表中,重复后覆盖，正好记录排在最后的位置，妙啊
	}

	return bc
}

//generate suffix and prefix array for pattern
// 模式串 [0, m-1]
// 子串 [0,m-2]
func generateGS(pattern string) ([]int, []bool) {
	m := len(pattern)
	suffix := make([]int, m)  //如果公共后缀子串的长度是 k，那我们就记录 suffix[k]=j（j 表示公共后缀子串的起始下标）
	prefix := make([]bool, m) //如果 j 等于 0，也就是说，公共后缀子串也是模式串的前缀子串，我们就记录 prefix[k]=true。

	//init
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	// 我们拿下标从 0 到 i 的子串（i 可以是 0 到 m-2）与整个模式串，求公共后缀子串。
	for i := 0; i < m-1; i++ {
		j := i // j来控制后缀子串长度
		k := 0
		//            从子串最后向前        从模式串最后向前
		for j >= 0 && pattern[j] == pattern[m-1-k] {
			j--
			k++
			suffix[k] = j + 1 //j--后要+1
		}

		if j == -1 { // 说明公共后缀子串也是模式串的前缀子串，j--了，从0变成了-1
			prefix[k] = true //方便在好后缀的后缀子串中，查找最长的、能跟模式串前缀子串匹配的后缀子串
		}
	}

	return suffix, prefix
}

//todo
func moveByGS(patternLength int, badCharStartIndex int, suffix []int, prefix []bool) int {

	//length of good suffix 好字符串的长度
	k := patternLength - badCharStartIndex - 1

	//complete match 说明模式串前缀有和好字符串相同的部分
	if suffix[k] != -1 {
		return badCharStartIndex + 1 - suffix[k] // 向后滑动几位
	}

	//partial match  没有完全匹配，要查找最长的、能跟模式串前缀子串匹配的好字符串的后缀子串
	for t := patternLength - 1; t > badCharStartIndex+1; t-- { // 从长到短遍历
		if prefix[t] {
			return t
		}
	}

	//no match 都不匹配，则移动整个模式串长度
	return patternLength

}

func bmSearch(main string, pattern string) int {
	//defensive
	if len(main) == 0 || len(pattern) == 0 || len(pattern) > len(main) {
		return -1
	}

	bc := generateBC(pattern)
	suffix, prefix := generateGS(pattern)

	n := len(main)
	m := len(pattern)

	// i : start index of main string i表示主串与模式串对齐的第一个字符下标
	step := 1                            // 向后移动几位
	for i := 0; i <= n-m; i = i + step { //根据之前得到的移动位数，向后移动
		subStr := main[i : i+m]                  //截取主串，开始比较
		k, j := findBadChar(subStr, pattern, bc) // 判断坏字符规则

		stepForBC := j - k // 相减得到坏字符情况下的后滑几位
		//j is bad char occur index
		if j == -1 {
			return i //没有坏字符，匹配成功，返回i
		}

		stepForGS := -1
		if j < m-1 {
			stepForGS = moveByGS(m, j, suffix, prefix) // 好后缀情况下的后滑几位，j是坏字符对应在模式串中的下标，m是模式串长度
		}

		//k is bad char index in pattern 比较好后缀、坏字符规则二者中的最大移动位数
		step = int(math.Max(float64(stepForBC), float64(stepForGS)))
	}

	return -1
}

func findBadChar(subStr string, pattern string, bc []int) (int, int) {

	j := -1
	k := -1
	badChar := rune(0)

	// 从后向前便利比较
	for index := len(subStr) - 1; index >= 0; index-- {
		if subStr[index] != pattern[index] {
			j = index                     // 坏字符对应在模式串中的下标
			badChar = rune(subStr[index]) // 坏字符
			break
		}
	}

	//if bad character exist, then find it's index at pattern
	if j > 0 {
		k = bc[int(badChar)] // 定位坏字符在模式串中最后下标
	}

	return k, j // j==-1，说明没有坏字符，匹配成功
}

func main() {

	main := "abcacabcbdcabcabcff"
	pattern := "cabcab"

	fmt.Println(bmSearch(main, pattern))
}
