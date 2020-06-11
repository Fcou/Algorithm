package ac

import (
	"fmt"
	"testing"
)

var trie *Trie

// Ac自动机借助Trie树结构，此为Trie节点
type AcNode struct {
	data         byte      // 存储字符
	children     []*AcNode // 字符集只包含a~z这26个字符
	isEndingChar bool      // 结尾字符为true
	length       int       // 当isEndingChar=true时，记录模式串长度
	fail         *AcNode   // 失败指针
}

func (ac *AcNode) SetData(data byte) {
	ac.data = data
}

func NewAcNode(d byte) *AcNode {
	return &AcNode{data: d, children: make([]*AcNode, 26), isEndingChar: false, length: -1, fail: nil}
}

// Trie树
type Trie struct {
	root *AcNode
}

// 创建Trie树的根节点
func NewTrie() {
	r := NewAcNode('/')
	trie = &Trie{root: r}
}

// 往Trie树中插入一个字符串
func (t *Trie) Insert(text []byte) {
	p := t.root
	for i := 0; i < len(text); i++ {
		index := BytesToInt(text[i]) - BytesToInt('a')
		if p.children[index] == nil {
			p.children[index] = NewAcNode(text[i])
		}
		p = p.children[index]
	}
	p.isEndingChar = true
	p.length = len(text)
	p.fail = nil
}

// 字符转化为int,笨办法
func BytesToInt(bys byte) int {
	m := map[byte]int{'a': 0, 'b': 1, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25}
	if v, ok := m[bys]; ok {
		return v
	}
	return -1
}

// 构建失败指针
func BuildFailurePointer() {
	queue := make([]*AcNode, 0)      // 队列,注意初始化大小要为0
	queue = append(queue, trie.root) // 根节点入队
	for len(queue) != 0 {
		p := queue[0]             // 取出首元素
		queue = queue[1:]         // 首元素出队
		for i := 0; i < 26; i++ { // 开始层次遍历
			pc := p.children[i]
			if pc == nil {
				continue
			}
			if p == trie.root {
				pc.fail = trie.root // root的孩子节点fail指向root
			} else { // 根据父节点的fail指针,来建立其孩子节点的fail指针
				q := p.fail
				for q != nil {
					qc := q.children[BytesToInt(pc.data)-BytesToInt('a')]
					if qc != nil {
						pc.fail = qc
						break
					}
					q = q.fail
				}
				if q == nil {
					pc.fail = trie.root
				}
			}
			queue = append(queue, pc) // 新节点入队
		}

	}
}

// 匹配，text是主串
func match(text []byte) {
	n := len(text)
	p := trie.root
	for i := 0; i < n; i++ {
		idx := BytesToInt(text[i]) - BytesToInt('a') //目前只能存储小写字母
		for p.children[idx] == nil && p != trie.root {
			p = p.fail // 失败指针发挥作用的地方
		}
		p = p.children[idx] // 模式树向后匹配
		if p == nil {
			p = trie.root // 如果没有匹配的，从root开始重新匹配
		}
		tmp := p
		for tmp != trie.root {
			if tmp.isEndingChar == true { // 打印出可以匹配的模式串
				pos := i - tmp.length + 1
				fmt.Printf("匹配起始下标 %d; 长度 %d\n", pos, tmp.length)
			}
			tmp = tmp.fail
		}
	}
}

func TestAc(t *testing.T) {
	NewTrie()
	trie.Insert([]byte("abcd"))
	trie.Insert([]byte("bcd"))
	trie.Insert([]byte("cd"))
	trie.Insert([]byte("d"))
	BuildFailurePointer()

	match([]byte("daddaddaceeecd"))
}
