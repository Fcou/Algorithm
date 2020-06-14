// 假设正则表达式中只包含“*”和“?”这两种通配符，并且对这两个通配符的语义稍微做些改变，
// 其中，“*”匹配任意多个（大于等于 0 个）任意字符，“?”匹配零个或者一个任意字符。
package regular_expression

import (
	"fmt"
	"testing"
)

type Pattern struct {
	matched bool
	pattern []byte // 正则表达式
	plen    int    // 正则表达式长度
}

func (p *Pattern) initPattern(pattern []byte, plen int) {
	p.matched = false
	p.pattern = pattern
	p.plen = plen
}

func (p *Pattern) match(text []byte, tlen int) bool { // 文本串及长度
	p.matched = false
	p.rmatch(0, 0, text, tlen)
	return p.matched
}

// ti表示目前匹配到文本串的位置， pj目前匹配到匹配串的位置
func (p *Pattern) rmatch(ti int, pj int, text []byte, tlen int) {
	if p.matched {
		return // 如果已经匹配了，就不要继续递归了
	}
	if pj == p.plen { // 正则表达式到结尾了
		if ti == tlen {
			p.matched = true // 文本串也到结尾了
		}
		return
	}
	if p.pattern[pj] == '*' { // *匹配任意个字符
		for k := 0; k <= tlen-ti; k++ {
			p.rmatch(ti+k, pj+1, text, tlen)
		}
	} else if p.pattern[pj] == '?' { // ?匹配0个或者1个字符
		p.rmatch(ti, pj+1, text, tlen)
		p.rmatch(ti+1, pj+1, text, tlen)
	} else if ti < tlen && p.pattern[pj] == text[ti] { // 纯字符匹配才行
		p.rmatch(ti+1, pj+1, text, tlen)
	}
}

func TestRregularExpreession(t *testing.T) {
	p := new(Pattern)
	p.initPattern([]byte("*a?"), 3)
	if p.match([]byte("xxayj"), 5) {
		fmt.Println("match.")
	} else {
		fmt.Println("don't match.")
	}
}
