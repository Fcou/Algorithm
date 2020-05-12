/*
栈是一种操作受限的数据结构，只支持入栈和出栈操作。后进先出是它最大的特点。
栈既可以通过数组实现，也可以通过链表来实现。
不管基于数组还是链表，入栈、出栈的时间复杂度都为 O(1)。


栈的应用场景：
1.浏览器前进后退功能（数组、双向链表也可以实现，就是记录一连串信息顺序而已）
我们使用两个栈，X 和 Y，我们把首次浏览的页面依次压入栈 X，当点击后退按钮时，
再依次从栈 X 中出栈，并将出栈的数据依次放入栈 Y。当我们点击前进按钮时，
我们依次从栈 Y 中取出数据，放入栈 X 中。当栈 X 中没有数据时，那就说明没有页面可以继续后退浏览了。
当栈 Y 中没有数据，那就说明没有页面可以点击前进按钮浏览了。
2.函数调用
3.表达式求值（两类信息，用两个容器存储处理，再构建二者之间的关系）
编译器就是通过两个栈来实现的。其中一个保存操作数的栈，另一个是保存运算符的栈。
我们从左向右遍历表达式，当遇到数字，我们就直接压入操作数栈；
当遇到运算符，就与运算符栈的栈顶元素进行比较。如果比运算符栈顶元素的优先级高，就将当前运算符压入栈；
如果比运算符栈顶元素的优先级低或者相同，从运算符栈中取栈顶运算符，从操作数栈的栈顶取 2 个操作数，
然后进行计算，再把计算完的结果压入操作数栈，继续比较。
4.栈在括号匹配中的应用（之前存入的信息，已对后来的信息有约束）
我们用栈来保存未匹配的左括号，从左到右依次扫描字符串。当扫描到左括号时，则将其压入栈中；
当扫描到右括号时，从栈顶取出一个左括号。如果能够匹配，比如“(”跟“)”匹配，“[”跟“]”匹配，“{”跟“}”匹配，
则继续扫描剩下的字符串。如果扫描的过程中，遇到不能配对的右括号，或者栈中没有数据，则说明为非法格式。
当所有的括号都扫描完成之后，如果栈为空，则说明字符串为合法格式；否则，说明有未匹配的左括号，为非法格式。

为什么函数调用要用“栈”来保存临时变量呢？用其他数据结构不行吗？
其实，我们不一定非要用栈来保存临时变量，只不过如果这个函数调用符合后进先出的特性，用栈这种数据结构来实现，
是最顺理成章的选择。
从调用函数进入被调用函数，对于数据来说，变化的是什么呢？
是作用域。所以根本上，只要能保证每进入一个新的函数，都是一个新的作用域就可以。
而要实现这个，用栈就非常方便。在进入被调用函数的时候，分配一段栈空间给这个函数的变量，
在函数结束的时候，将栈顶复位，正好回到调用函数的作用域内。
*/
//一个可扩展栈，基于切片（可自动扩展）
package stack

type myStack struct {
	items []string
	count int //当前栈内元素个数
}

func (this *myStack) newStack(n int) *myStack {
	this.items = make([]string, n)
	this.count = 0
	return this
}

func (this *myStack) push(item string) {
	this.items = append(this.items, item)
	this.count++
	return
}

func (this *myStack) pop() (result string, ok bool) {
	if this.count < 1 {
		result = ""
		ok = false
		return
	}
	result = this.items[this.count-1]
	ok = true
	this.count--
	this.items = this.items[:this.count-1]
	return
}
