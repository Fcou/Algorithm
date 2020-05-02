/* 给出一个字符串数组 words 组成的一本英语字典。
从中找出最长的一个单词，该单词是由words字典中其他单词逐步添加一个字母组成。
若其中有多个可行的答案，则返回答案中字典字母顺序最前的单词
字符串都只包含小写字母
words最多包含1000个单词
words单词最多30个字母*/

//有时直接遍历效率低，先把信息组织好，然后再查找
package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]bool, len(words))

	res := "a" //最初最长的单词初始化设置为”a"是个好选择？不应该是words[0]吧

	for _, w := range words {
		n := len(w)
		if n == 1 {
			m[w] = true
		} else if m[w[:n-1]] {

			m[w] = true
			if len(res) < len(w) {
				res = w
			}
		}
	}

	return res
}

func main() {
	s := []string{"go", "bravo", "gopher", "alone", "g", "gos", "gelta"}
	sort.Strings(s)
	fmt.Println(s)
	longs := longestWord(s)
	fmt.Println("The logest word is: ", longs)
}

//这种算法感觉没用到字典树
