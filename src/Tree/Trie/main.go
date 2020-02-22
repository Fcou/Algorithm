/*
Trie树，即字典树，又称单词查找树或键树
核心思想就是通过最大限度地减少无谓的字符串比较，使得查询高效率，即「用空间换时间」，再利用共同前缀来提高查询效率。
*/

package main

import "fmt"

// 字典树节点
type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode), isEnd: false}
}

// 字典树
type Trie struct {
	root *TrieNode
}

// 构造字典树
func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// 向字典树中插入一个单词
func (trie *Trie) Insert(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

// 搜索字典树中是否存在指定单词
func (trie *Trie) Search(word string) bool {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			return false
		}
		node = node.children[word[i]]
	}
	return node.isEnd
}

// 判断字典树中是否有指定前缀的单词
func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("hello")

	ok := trie.Search("hollo")
	if ok {
		fmt.Println("存在指定单词")
	} else {
		fmt.Println("不存在指定单词")
	}

	ok = trie.StartsWith("hel")
	if ok {
		fmt.Println("存在指定前缀")
	} else {
		fmt.Println("不存在指定前缀")
	}

}
