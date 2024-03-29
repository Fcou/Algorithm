// 我们可以把物品依次排列，整个问题就分解为了 n 个阶段，每个阶段对应一个物品怎么选择。
// 先对第一个物品进行处理，选择装进去或者不装进去，然后再递归地处理剩下的物品。
package bp


var maxW int= -1 //存储背包中物品总重量的最大值
// cw表示当前已经装进去的物品的重量和；i表示考察到哪个物品了；
// w背包重量；items表示每个物品的重量；n表示物品个数
// 假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
// f(0, 0, a, 10, 100)

var weight []int  = []int{2，2，4，6，3}  // 物品重量
var  n int = 5 // 物品个数
var  w int  = 9// 背包承受的最大重量
var  mem [5][10]bool    // 备忘录，默认值false
// 创建二维数组例子
// moved := make([][]bool, m)
// for i := 0; i < m; i++ {
//   moved[i] = make([]bool, n)
// }


func f(i, cw int, items int[] , n, w int ) {
  if (cw == w || i == n) { // cw==w表示装满了;i==n表示已经考察完所有的物品
    if cw > maxW {
		  maxW = cw
	  }
    return
  }

  if mem[i][cw] {
    return // 重复状态
  }
  mem[i][cw] = true // 不重复，则记录(i, cw)这个状态

  f(i+1, cw, items, n, w) // 不装，考察下一项

  if cw + items[i] <= w {// 已经超过可以背包承受的重量的时候，就不要再装了，剪枝
    f(i+1,cw + items[i], items, n, w) // 装，考察下一项
  }
}