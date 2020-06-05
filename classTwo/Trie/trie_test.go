package trie

import "testing"

type TrieNode struct {
	data         byte
	children     []*TrieNode // 这部分可以替代成跳表、红黑树等，时间换空间
	isEndingChar bool
}

func NewTrieNode(b byte, end bool) *TrieNode {
	return &TrieNode{data: b, children: make([]*TrieNode, 26), isEndingChar: end}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	r := NewTrieNode('/', false)
	return &Trie{root: r}
}

// 往Trie树中插入一个字符串
func (t *Trie) Insert(text []byte) {
	p := t.root
	for i := 0; i < len(text); i++ {
		index := BytesToInt(text[i]) - BytesToInt('a')
		if p.children[index] == nil {
			p.children[index] = NewTrieNode(text[i], false)
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

// 在Trie树中查找一个字符串
func (t *Trie) Find(pattern []byte) bool {
	p := t.root
	for i := 0; i < len(pattern); i++ {
		index := BytesToInt(pattern[i]) - BytesToInt('a') //目前只能存储小写字母
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	if p.isEndingChar == false {
		return false // 不能完全匹配，只是前缀
	}
	return true // 找到pattern
}

// 笨办法
func BytesToInt(bys byte) int {
	m := map[byte]int{'a': 0, 'b': 1, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25}
	if v, ok := m[bys]; ok {
		return v
	}
	return -1
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert([]byte("hello"))
	trie.Insert([]byte("world"))
	t.Log(trie.Find([]byte("hello")))
	t.Log(trie.Find([]byte("hel")))
	/*
	   === RUN   TestTrie
	   --- PASS: TestTrie (0.00s)
	       trie_test.go:65: true
	       trie_test.go:66: false
	   PASS
	   ok      command-line-arguments  0.002s
	*/
}
