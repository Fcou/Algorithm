package main

import (
	"fmt"
	"math"
)

//bc: pattern char index hash mapping  坏字符规则，快速定位
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
func generateGS(pattern string) ([]int, []bool) {
	m := len(pattern)
	suffix := make([]int, m)
	prefix := make([]bool, m)

	//init
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	// 依次匹配后缀子串在模式串中的重复开始下标，长度从 1 到 m-1（好字符长度最多是m-1,如果是m说明已经匹配，坏字符已考虑此情况）
	for i := 0; i < m-1; i++ {
		j := i // j来控制后缀子串长度
		k := 0
		//            从下标0开始向          从最后向前
		for j >= 0 && pattern[j] == pattern[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}

		if j == -1 { // 说明公共后缀子串也是模式串的前缀子串
			prefix[k] = true //方便在好后缀的后缀子串中，查找最长的、能跟模式串前缀子串匹配的后缀子串
		}
	}

	return suffix, prefix
}

//todo
func moveByGS(patternLength int, badCharStartIndex int, suffix []int, prefix []bool) int {

	//length of good suffix 好字符串的长度
	k := patternLength - badCharStartIndex - 1

	//complete match 说明完全匹配成功
	if suffix[k] != -1 {
		return badCharStartIndex + 1 - suffix[k] // 向后滑动几位
	}

	//partial match  没有完全匹配，要查找最长的、能跟模式串前缀子串匹配的后缀子串
	for t := patternLength - 1; t > badCharStartIndex+1; t-- { // 从长到短遍历
		if prefix[t] {
			return t
		}
	}

	//no match
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
	step := 1 // 向后移动几位
	for i := 0; i <= n-m; i = i + step {
		subStr := main[i : i+m]                  //截取主串，开始比较
		k, j := findBadChar(subStr, pattern, bc) // 判断坏字符规则

		stepForBC := j - k // 相减得到坏字符情况下的后滑几位
		//j is bad char occur index
		if j == -1 {
			return i //匹配成功，返回i
		}

		stepForGS := -1
		if j < m-1 {
			stepForGS = moveByGS(m, j, suffix, prefix)
		}

		//k is bad char index in pattern 比较好、坏字符规则最大移动位数
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
